package migrator

import (
	"fmt"
	"github.com/99neerajsharma/FamTube/internal/model"
	"gorm.io/gorm"
)

func Migrate(pg *gorm.DB) error {

	err := pg.AutoMigrate(&model.Video{})
	if err != nil {
		fmt.Println(err)
	}
	return err
}
