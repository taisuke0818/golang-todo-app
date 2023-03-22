package todo

import (
	"context"

	"api.todo/internal/config"
	"api.todo/internal/store"
	todopb "api.todo/protobuf/todo/v1"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	todopb.UnimplementedTodoServiceServer
	todopb.TodoServiceClient
	storeClient *store.Client
	Config      *config.Config
}

func (s *Service) Init(ctx context.Context) error {
	if client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(s.Config.MongoUri).
		SetAuth(options.Credential{
			AuthSource: s.Config.MongoDatabase,
			Username:   s.Config.MongoUsername,
			Password:   s.Config.MongoPassword,
		})); err != nil {
		return err
	} else {
		s.storeClient = store.New(client)
	}

	return nil
}
