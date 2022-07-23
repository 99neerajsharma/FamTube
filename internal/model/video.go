package model

import (
	"time"
)

type Video struct {
	ID                  string    `gorm:"type:varchar;size:12;primaryKey;index;not null" json:"id"`
	CreatedAt           time.Time `gorm:"autoCreateTime;not null;default:now();" json:"created_at"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime;not null;default:now();" json:"updated_at"`
	Title               string    `gorm:"type:varchar;not null;index" json:"title"`
	Description         string    `gorm:"type:varchar;default:null;index" json:"description"`
	DefaultThumbnailURL string    `gorm:"type:varchar;default:null" json:"default_thumbnail_url"`
	MediumThumbnailURL  string    `gorm:"type:varchar;default:null" json:"medium_thumbnail_url"`
	HighThumbnailURL    string    `gorm:"type:varchar;default:null" json:"high_thumbnail_url"`
	ChannelName         string    `gorm:"type:varchar;not null;" json:"channel_name"`
	PublishedAt         time.Time `gorm:"not null;index" json:"published_at"`
}

func (v *Video) TableName() string {
	return "video"
}
