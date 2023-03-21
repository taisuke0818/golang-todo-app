package codec

import (
	"fmt"
	"strings"
	"unicode"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func MarshalMessage(m proto.Message) bson.D {
	var l = []bson.E{}
	rv := m.ProtoReflect()
	fds := m.ProtoReflect().Descriptor().Fields()
	// Range over all known fields.
	// rv.Range()はゼロ値は無視されてしまうため使わない
	for i := 0; i < fds.Len(); i++ {
		fd := fds.Get(i)
		v := rv.Get(fd)
		if rv.Has(fd) {
			switch fd.Kind() {
			case protoreflect.BoolKind:
				l = append(l, bson.E{Key: string(fd.JSONName()), Value: v.Bool()})
			case protoreflect.Int32Kind, protoreflect.Int64Kind,
				protoreflect.Sint32Kind, protoreflect.Sint64Kind,
				protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
				l = append(l, bson.E{Key: string(fd.JSONName()), Value: v.Int()})
			case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
				protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
				l = append(l, bson.E{Key: string(fd.JSONName()), Value: v.Uint()})
			case protoreflect.FloatKind, protoreflect.DoubleKind:
				l = append(l, bson.E{Key: string(fd.JSONName()), Value: v.Float()})
			case protoreflect.StringKind:
				l = append(l, bson.E{Key: string(fd.JSONName()), Value: v.String()})
			case protoreflect.BytesKind:
				l = append(l, bson.E{Key: string(fd.JSONName()), Value: v.Bytes()})
			default:
				l = append(l, bson.E{Key: string(fd.JSONName()), Value: v.Interface()})
			}
		} else {
			l = append(l, bson.E{Key: string(fd.JSONName()), Value: zeroValue(fd, v)})
		}
	}
	return l
}

func DecodeMessage(data bson.Raw, dst protoreflect.ProtoMessage) error {
	// ObjectId を取得
	// https://www.mongodb.com/docs/manual/core/document/#the-_id-field
	primaryKey := data.Lookup("_id")
	if primaryKey.Type != bsontype.ObjectID {
		return fmt.Errorf("object id must not be empty")
	}
	rv := dst.ProtoReflect()
	md := rv.Descriptor()
	fds := md.Fields()
	pkName := GetPrimaryKeyName(string(md.FullName()))
	if fd := fds.ByName(protoreflect.Name(pkName)); fd != nil {
		rv.Set(fd, protoreflect.ValueOfString(primaryKey.ObjectID().Hex()))
	}
	return nil
}

func GetPrimaryKeyValue(m protoreflect.Message) (string, error) {
	pkName := GetPrimaryKeyName(string(m.Descriptor().FullName()))
	fd := m.Descriptor().Fields().ByName(protoreflect.Name(pkName))
	if fd == nil {
		return "", fmt.Errorf("fail to get primary key from message: pkName=%s", pkName)
	}
	return m.Get(fd).String(), nil
}

func GetPrimaryKeyName(typename string) string {
	if typename == "" {
		return ""
	}
	typeslice := strings.Split(typename, ".")
	t := typeslice[len(typeslice)-1]
	idname := strings.ToLower(t[:1]) + t[1:] + "_id"
	return toSnakeCase(idname)
}

func toSnakeCase(s string) string {
	b := &strings.Builder{}
	for i, r := range s {
		if i == 0 {
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			b.WriteRune('_')
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
