package controller

import (
	"github.com/99neerajsharma/FamTube/internal/api/video"
	"github.com/99neerajsharma/FamTube/internal/service"
)

type IAPIController interface {
}

type APIController struct {
	VideoController video.IVideoAPIController
}

func APIControllerInitializer(videoService *service.VideoService) *APIController {
	return &APIController{
		VideoController: &video.VideoAPIController{VideoService: videoService},
	}
}
