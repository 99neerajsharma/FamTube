package video

import (
	"github.com/99neerajsharma/FamTube/internal/service"
	"github.com/gin-gonic/gin"
)

type IVideoAPIController interface {
	GetVideosV1APIHandler(c *gin.Context)
	GetSearchVideosV1APIHandler(c *gin.Context)
}

type VideoAPIController struct {
	VideoService service.IVideoService
}
