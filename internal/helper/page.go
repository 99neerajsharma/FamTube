package helper

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetOffsetAndPageSize(c *gin.Context) (int, int) {
	offset := 0
	pageSize := 10
	queryParam := c.Request.URL.Query()
	if pS, ok := queryParam["pageSize"]; ok {
		if len(pS) > 0 {
			intVal, err := strconv.Atoi(pS[0])
			if err == nil && intVal > 5 && intVal < 50 {
				pageSize = intVal
			}
		}
	}
	if page, ok := queryParam["page"]; ok {
		if len(page) > 0 {
			intVal, err := strconv.Atoi(page[0])
			if err == nil && intVal > 0 {
				offset = pageSize * (intVal - 1)
			}
		}
	}

	return offset, pageSize
}
