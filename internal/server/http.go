package server

import (
	"context"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/config"
	"go.uber.org/fx"
	"log"
)

func NewHttpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func RegisterServerHooks(lifecycle fx.Lifecycle, router *gin.Engine, configYAML *config.YAML) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				log.Println("starting server")
				go endless.ListenAndServe(fmt.Sprintf(":%v", configYAML.Get("server.port")), router)
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		})
}
