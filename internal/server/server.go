package server

import (
	"context"
	"github.com/99neerajsharma/FamTube/internal/db"
	"github.com/99neerajsharma/FamTube/internal/migrator"
	"go.uber.org/fx"
)

func serverFxApp(ctx context.Context) *fx.App {
	return fx.New(
		fx.Provide(func() context.Context { return ctx }),
		fx.Provide(NewHttpRouter),
		fx.Provide(db.PostgresInitializer),
		fx.Invoke(migrator.Migrate),
		fx.Invoke(RegisterServerHooks),
		fx.Invoke(registerRoutes),
	)
}

func Run(ctx context.Context) {
	serverFxApp(ctx).Run()
}
