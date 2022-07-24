package server

import (
	"github.com/99neerajsharma/FamTube/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func internalAPIRoute(apiRouter *gin.RouterGroup, apiController *controller.APIController) {
	apiRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"server_status": "running"})
	})

	apiRouter.GET("/api/v1/videos", apiController.VideoController.GetVideosV1APIHandler)
	apiRouter.GET("/api/v1/videos/search", apiController.VideoController.GetSearchVideosV1APIHandler)
}

func registerRoutes(router *gin.Engine, apiController *controller.APIController) {
	apiRouterGroup := router.Group("")

	internalAPIRoute(apiRouterGroup, apiController)
}
