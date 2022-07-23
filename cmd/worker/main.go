package main

import (
	"github.com/99neerajsharma/FamTube/app"
	"log"
	"os"
)

func main() {
	workerApp := app.WorkerApp()
	if err := workerApp.Run(os.Args); err != nil {
		log.Println(err)
	}
}
