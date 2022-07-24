package service

import (
	"github.com/99neerajsharma/FamTube/internal/model"
	"gorm.io/gorm"
	"log"
	"strings"
)

type IVideoService interface {
	AllVideosPaginated(offset int, pageSize int) []*model.Video
	SearchVideosPaginated(offset int, pageSize int, query string) []*model.Video
}

type VideoService struct {
	pg *gorm.DB
}

func VideoServiceInitializer(pg *gorm.DB) *VideoService {
	return &VideoService{pg: pg}
}

func (vS *VideoService) AllVideosPaginated(offset int, pageSize int) []*model.Video {
	var videos []*model.Video

	if err := vS.pg.Model(&model.Video{}).Order("published_at desc").Offset(offset).Limit(pageSize).
		Find(&videos).Error; err != nil {
		log.Println(err)
	}
	return videos
}

func (vS *VideoService) SearchVideosPaginated(offset int, pageSize int, query string) []*model.Video {
	var videos []*model.Video

	query = strings.ToLower(query)
	query = "%" + query + "%"
	if err := vS.pg.Model(&model.Video{}).Where("lower(title) LIKE ? OR lower(description) LIKE ?", query, query).
		Offset(offset).Limit(pageSize).Order("published_at desc").
		Find(&videos).Error; err != nil {
		log.Println(err)
	}
	return videos
}
