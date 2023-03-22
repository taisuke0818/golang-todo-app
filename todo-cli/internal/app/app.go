package app

import (
	"context"
	"sync"

	todopb "api.todo/protobuf/todo/v1"
	"api.todo/todo-cli/internal/grpc/dialer"
)

type App struct {
	URL string

	initOnce sync.Once

	initErr error

	Client todopb.TodoServiceClient
}

func (app *App) Init(ctx context.Context) error {
	app.initOnce.Do(func() {
		conn, err := dialer.DialContext(ctx, app.URL)
		if err != nil {
			app.initErr = err
			return
		}
		app.Client = todopb.NewTodoServiceClient(conn)
	})
	return app.initErr
}
