package app

import (
	"context"
)

type ctxKey string

const (
	ctxKeyApp = ctxKey("app")
)

func AppFromContext(ctx context.Context) *App {
	return ctx.Value(ctxKeyApp).(*App)
}

func ContextWithApp(parent context.Context, v *App) context.Context {
	return context.WithValue(parent, ctxKeyApp, v)
}
