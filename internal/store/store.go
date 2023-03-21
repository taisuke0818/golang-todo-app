package store

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"api.todo/internal"
	"api.todo/internal/store/codec"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (c *Client) CreateMessage(ctx context.Context, db string, m proto.Message) error {
	rv := m.ProtoReflect()
	md := rv.Descriptor()
	fds := md.Fields()

	// 作成時の作成/更新日時は同じ
	createdAt := internal.TimeNow().UnixMicro()
	if fd := fds.ByName(protoreflect.Name("created_at")); fd != nil {
		rv.Set(fd, protoreflect.ValueOfInt64(createdAt))
	}
	if fd := fds.ByName(protoreflect.Name("updated_at")); fd != nil {
		rv.Set(fd, protoreflect.ValueOfInt64(createdAt))
	}
	if fd := fds.ByName(protoreflect.Name("deleted_at")); fd != nil {
		rv.Set(fd, protoreflect.ValueOfInt64(0))
	}
	col := c.c.Database(db).Collection(string(proto.MessageName(m)))
	res, err := col.InsertOne(ctx, codec.MarshalMessage(m))
	if err != nil {
		return err
	}
	// ObjectId を設定
	if v, ok := res.InsertedID.(primitive.ObjectID); ok {
		pkName := codec.GetPrimaryKeyName(string(md.FullName()))
		if fd := fds.ByName(protoreflect.Name(pkName)); fd != nil {
			rv.Set(fd, protoreflect.ValueOfString(v.Hex()))
		}
	}
	return nil
}

func (c *Client) GetMessage(ctx context.Context, db string, m proto.Message) error {
	rv := m.ProtoReflect()

	hex, err := codec.GetPrimaryKeyValue(rv)
	if err != nil {
		return err
	}
	objectId, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return err
	}
	col := c.c.Database(db).Collection(string(proto.MessageName(m)))
	res := col.FindOne(ctx, &bson.M{"_id": objectId})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return ErrNotFound
	}
	if err := res.Decode(m); err != nil {
		return err
	}
	current, err := res.DecodeBytes()
	if err != nil {
		return err
	}
	if err = codec.DecodeMessage(current, m); err != nil {
		return err
	}
	return nil
}

func (c *Client) ListMessage(ctx context.Context, db string, out any) (err error) {
	if reflect.TypeOf(out).Kind() != reflect.Pointer {
		panic(fmt.Sprintf("out must be a pointer to slice: %T", out))
	}
	m := assumeElementMessage(out)

	col := c.c.Database(db).Collection(string(proto.MessageName(m)))
	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer func() {
		err = cursor.Close(ctx)
	}()
	rv := reflect.Indirect(reflect.ValueOf(out))
	for cursor.Next(ctx) {
		dst := proto.Clone(m)
		if err := cursor.Decode(dst); err != nil {
			return err
		}
		if err = codec.DecodeMessage(cursor.Current, dst); err != nil {
			return err
		}
		// 論理削除されていればスキップ
		dstrv := dst.ProtoReflect()
		if fd := dstrv.Descriptor().Fields().ByName(protoreflect.Name("deleted_at")); fd != nil && fd.Kind() == protoreflect.Int64Kind {
			if dstrv.Get(fd).Int() > 0 {
				continue
			}
		}
		rv.Set(reflect.Append(rv, reflect.ValueOf(dst)))
	}
	return nil
}

func (c *Client) UpdateMessage(ctx context.Context, db string, m proto.Message) error {
	rv := m.ProtoReflect()
	md := rv.Descriptor()
	fds := md.Fields()

	dst := proto.Clone(m)
	if err := c.GetMessage(ctx, db, dst); err != nil {
		return err
	}
	cur := dst.ProtoReflect()

	fd := fds.ByName(protoreflect.Name("updated_at"))
	if fd != nil {
		if fd.Kind() == protoreflect.Int64Kind && rv.Get(fd).Int() != cur.Get(fd).Int() {
			// 更新日時が一致しなければ、更新不可（排他処理）
			return FailedPrecondition
		}
		// 更新日時を更新
		rv.Set(fd, protoreflect.ValueOfInt64(internal.TimeNow().UnixMicro()))
	}
	if fd := fds.ByName(protoreflect.Name("created_at")); fd != nil && fd.Kind() == protoreflect.Int64Kind {
		rv.Set(fd, protoreflect.ValueOfInt64(cur.Get(fd).Int()))
	}
	if fd := fds.ByName(protoreflect.Name("deleted_at")); fd != nil && fd.Kind() == protoreflect.Int64Kind {
		rv.Set(fd, protoreflect.ValueOfInt64(cur.Get(fd).Int()))
	}
	hex, err := codec.GetPrimaryKeyValue(rv)
	if err != nil {
		return err
	}
	objectId, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return err
	}
	col := c.c.Database(db).Collection(string(proto.MessageName(m)))
	res := col.FindOneAndUpdate(ctx, &bson.M{"_id": objectId},
		bson.D{{
			Key:   "$set",
			Value: codec.MarshalMessage(m),
		}},
		options.FindOneAndUpdate().
			SetUpsert(false).
			SetReturnDocument(options.After),
	)
	if err := res.Decode(m); err != nil {
		return err
	}
	current, err := res.DecodeBytes()
	if err != nil {
		return err
	}
	if err = codec.DecodeMessage(current, m); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteMessage(ctx context.Context, db string, m proto.Message) error {
	rv := m.ProtoReflect()
	md := rv.Descriptor()
	fds := md.Fields()

	if fd := fds.ByName(protoreflect.Name("deleted_at")); fd != nil {
		// 削除日時を更新
		rv.Set(fd, protoreflect.ValueOfInt64(internal.TimeNow().UnixMicro()))
	}

	hex, err := codec.GetPrimaryKeyValue(rv)
	if err != nil {
		return err
	}
	objectId, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return err
	}
	col := c.c.Database(db).Collection(string(proto.MessageName(m)))
	// 論理削除
	res, err := col.UpdateByID(ctx, objectId,
		bson.D{{
			Key: "$set",
			Value: bson.D{{
				Key:   "deletedat",
				Value: internal.TimeNow().UnixMicro()},
			}}},
		options.Update().SetUpsert(false),
	)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return ErrNotFound
	}
	return nil
}

func assumeElementMessage(l any) proto.Message {
	// *[]*pb.X or []*pb.X
	typ := reflect.TypeOf(l)
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Slice {
		panic(fmt.Sprintf("l must be a pointer to slice or slice: %T", l))
	}
	// *pb.X
	typ = typ.Elem()
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	m, ok := reflect.New(typ).Interface().(proto.Message)
	if !ok {
		panic(fmt.Sprintf("l must be a pointer to slice or slice of proto.Message: %T", l))
	}
	return m
}
