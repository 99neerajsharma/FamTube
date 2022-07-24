package server

import (
	"context"
	"github.com/99neerajsharma/FamTube/internal/config"
	"github.com/99neerajsharma/FamTube/internal/controller"
	"github.com/99neerajsharma/FamTube/internal/db"
	"github.com/99neerajsharma/FamTube/internal/migrator"
	"github.com/99neerajsharma/FamTube/internal/service"
	"go.uber.org/fx"
)

func serverFxApp(ctx context.Context) *fx.App {
	return fx.New(
		fx.Provide(func() context.Context { return ctx }),
		fx.Provide(NewHttpRouter),
		fx.Provide(db.PostgresInitializer),
		fx.Provide(service.VideoServiceInitializer),
		fx.Provide(config.ConfigInitializer),
		fx.Provide(controller.APIControllerInitializer),
		fx.Invoke(migrator.Migrate),
		fx.Invoke(RegisterServerHooks),
		fx.Invoke(registerRoutes),
	)
}

func Run(ctx context.Context) {
	serverFxApp(ctx).Run()
}
