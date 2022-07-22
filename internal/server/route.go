package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func internalAPIRoute(apiRouter *gin.RouterGroup) {
	apiRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"server_status": "running"})
	})
}

func registerRoutes(router *gin.Engine) {
	apiRouterGroup := router.Group("")

	internalAPIRoute(apiRouterGroup)
}
