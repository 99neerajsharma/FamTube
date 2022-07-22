package main

import (
	"github.com/99neerajsharma/FamTube/app"
	"log"
	"os"
)

func main() {
	serverApp := app.ServerApp()
	if err := serverApp.Run(os.Args); err != nil {
		log.Println(err)
	}
}
