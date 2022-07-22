package server

import (
	"context"
	"go.uber.org/fx"
)

func serverFxApp(ctx context.Context) *fx.App {
	return fx.New(
		fx.Provide(func() context.Context { return ctx }),
		fx.Provide(NewHttpRouter),
		fx.Invoke(RegisterServerHooks),
		fx.Invoke(registerRoutes),
	)
}

func Run(ctx context.Context) {
	serverFxApp(ctx).Run()
}
