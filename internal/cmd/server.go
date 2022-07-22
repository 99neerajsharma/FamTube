package cmd

import (
	"github.com/99neerajsharma/FamTube/internal/server"
	"github.com/urfave/cli/v2"
)

var SeverStartCommand = &cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Starts HTTP server",
	Action:  startAction,
}

func startAction(ctx *cli.Context) (err error) {
	server.Run(ctx.Context)
	return nil
}
