package video

import (
	"github.com/99neerajsharma/FamTube/internal/helper"
	"github.com/gin-gonic/gin"
)

func (API *VideoAPIController) GetVideosV1APIHandler(c *gin.Context) {
	offset, pageSize := helper.GetOffsetAndPageSize(c)
	videos := API.VideoService.AllVideosPaginated(offset, pageSize)
	resp := make(map[string]interface{})
	resp["videos"] = videos
	c.Header("Content-Type", "application/json")
	c.JSON(200, resp)
}

func (API *VideoAPIController) GetSearchVideosV1APIHandler(c *gin.Context) {
	queryParam := c.Request.URL.Query()
	query := ""
	if val, ok := queryParam["query"]; ok {
		query = val[0]
	}
	offset, pageSize := helper.GetOffsetAndPageSize(c)
	videos := API.VideoService.SearchVideosPaginated(offset, pageSize, query)
	resp := make(map[string]interface{})
	resp["videos"] = videos
	c.Header("Content-Type", "application/json")
	c.JSON(200, resp)
}
