package cmd

import (
	"github.com/99neerajsharma/FamTube/internal/worker"
	"github.com/urfave/cli/v2"
)

var WorkerStartCommand = &cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Starts Worker",
	Action:  workerStartAction,
}

func workerStartAction(ctx *cli.Context) (err error) {
	worker.Run(ctx.Context)
	return nil
}
