package todo

import (
	"context"
	"errors"

	"api.todo/internal/store"
	todopb "api.todo/protobuf/todo/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CreateTodoTask(ctx context.Context, req *todopb.CreateTodoTaskRequest) (*todopb.CreateTodoTaskResponse, error) {
	logger := ctxlogrus.Extract(ctx)

	todo := req.TodoTask
	if err := s.storeClient.CreateMessage(ctx, s.Config.MongoDatabase, todo); err != nil {
		logger.WithError(err).Error("fail to CreateTodoTask")
		return nil, status.Error(codes.Internal, "Internal")
	}
	out := new(todopb.CreateTodoTaskResponse)
	out.TodoTask = todo
	return out, nil
}

func (s *Service) ListTodoTasks(ctx context.Context, req *todopb.ListTodoTasksRequest) (*todopb.ListTodoTasksResponse, error) {
	logger := ctxlogrus.Extract(ctx)

	var todos []*todopb.TodoTask
	if err := s.storeClient.ListMessage(ctx, s.Config.MongoDatabase, &todos); err != nil {
		logger.WithError(err).Error("fail to ListTodoTasks")
		return nil, status.Error(codes.Internal, "Internal")
	}
	out := new(todopb.ListTodoTasksResponse)
	out.Items = todos
	out.Total = int64(len(todos))
	return out, nil
}

func (s *Service) GetTodoTask(ctx context.Context, req *todopb.GetTodoTaskRequest) (*todopb.GetTodoTaskResponse, error) {
	logger := ctxlogrus.Extract(ctx)

	task := todopb.TodoTask{TodoTaskId: req.TodoTaskId}
	if err := s.storeClient.GetMessage(ctx, s.Config.MongoDatabase, &task); err != nil {
		logger.WithError(err).Error("fail to GetMessage")
		switch {
		case errors.Is(err, store.ErrNotFound):
			return nil, status.Error(codes.NotFound, "NotFound")
		default:
			return nil, status.Error(codes.Internal, "Internal")
		}
	}
	out := new(todopb.GetTodoTaskResponse)
	out.TodoTask = &task
	return out, nil
}

func (s *Service) UpdateTodoTask(ctx context.Context, req *todopb.UpdateTodoTaskRequest) (*todopb.UpdateTodoTaskResponse, error) {
	logger := ctxlogrus.Extract(ctx)

	todo := req.TodoTask
	if err := s.storeClient.UpdateMessage(ctx, s.Config.MongoDatabase, todo); err != nil {
		logger.WithError(err).Error("fail to UpdateMessage")
		switch {
		case errors.Is(err, store.ErrNotFound):
			return nil, status.Error(codes.NotFound, "NotFound")
		case errors.Is(err, store.FailedPrecondition):
			return nil, status.Error(codes.FailedPrecondition, "FailedPrecondition")
		default:
			return nil, status.Error(codes.Internal, "Internal")
		}
	}
	out := new(todopb.UpdateTodoTaskResponse)
	out.TodoTask = todo
	return out, nil
}

func (s *Service) DeleteTodoTask(ctx context.Context, req *todopb.DeleteTodoTaskRequest) (*emptypb.Empty, error) {
	logger := ctxlogrus.Extract(ctx)

	todo := req.TodoTask
	if err := s.storeClient.DeleteMessage(ctx, s.Config.MongoDatabase, todo); err != nil {
		logger.WithError(err).Error("fail to DeleteMessage")
		switch {
		case errors.Is(err, store.ErrNotFound):
			return nil, status.Error(codes.NotFound, "NotFound")
		default:
			return nil, status.Error(codes.Internal, "Internal")
		}
	}
	return new(emptypb.Empty), nil
}
