package app

import (
	"github.com/99neerajsharma/FamTube/internal/cmd"
	"github.com/urfave/cli/v2"
)

func ServerApp() *cli.App {
	app := cli.NewApp()
	app.Name = "FamTube Server"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		cmd.SeverStartCommand,
	}

	return app
}
