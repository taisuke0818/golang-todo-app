package main

import (
	"context"
	"os"

	todo "api.todo"
	"api.todo/internal/config"
	startup "api.todo/internal/startup"
	todopb "api.todo/protobuf/todo/v1"
	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func newServer(ctx context.Context) (*grpc.Server, error) {
	logger := ctxlogrus.Extract(ctx)

	var hdlr todo.Service
	if c, err := config.LoadConfig(); err != nil {
		logger.WithError(err).Error("fail to load config")
		return nil, err
	} else {
		hdlr.Config = c
	}
	if err := hdlr.Init(ctx); err != nil {
		logger.WithError(err).Error("fail to init service")
		return nil, err
	}
	srv := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			grpclogrus.StreamServerInterceptor(logrus.NewEntry(logrus.StandardLogger())),
			grpcvalidator.StreamServerInterceptor(),
		),
		grpc.ChainUnaryInterceptor(
			grpclogrus.UnaryServerInterceptor(logrus.NewEntry(logrus.StandardLogger())),
			grpcvalidator.UnaryServerInterceptor(),
		),
	)
	todopb.RegisterTodoServiceServer(srv, &hdlr)
	return srv, nil
}

func main() {
	os.Exit(startup.RunGRPCServer(&startup.GRPCServerConfig{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		New:    newServer,
	}))
}
