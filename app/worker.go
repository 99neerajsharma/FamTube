package app

import (
	"github.com/99neerajsharma/FamTube/internal/cmd"
	"github.com/urfave/cli/v2"
)

func WorkerApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Async worker"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		cmd.WorkerStartCommand,
	}

	return app
}
