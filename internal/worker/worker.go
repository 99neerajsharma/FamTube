package worker

import (
	"context"
	"github.com/99neerajsharma/FamTube/internal/db"
	"github.com/99neerajsharma/FamTube/internal/migrator"
	"go.uber.org/fx"
)

func workerFxApp(ctx context.Context) *fx.App {
	return fx.New(
		fx.Provide(func() context.Context { return ctx }),
		fx.Provide(db.PostgresInitializer),
		fx.Invoke(migrator.Migrate),
		fx.Invoke(SeedVideoData),
	)
}

func Run(ctx context.Context) {
	workerFxApp(ctx).Run()
}
