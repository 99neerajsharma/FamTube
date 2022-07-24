package migrator

import (
	"github.com/99neerajsharma/FamTube/internal/model"
	"gorm.io/gorm"
	"log"
)

func Migrate(pg *gorm.DB) error {

	err := pg.AutoMigrate(&model.Video{})
	if err != nil {
		log.Println(err)
	}
	return nil
}
