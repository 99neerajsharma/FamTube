package db

import (
	"fmt"
	"go.uber.org/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func getPostgresURL(configYAML *config.YAML) string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		configYAML.Get("db.pg.host"), configYAML.Get("db.pg.user"), configYAML.Get("db.pg.password"),
		configYAML.Get("db.pg.name"), configYAML.Get("db.pg.port"))
}
func PostgresInitializer(configYAML *config.YAML) *gorm.DB {
	db, err := gorm.Open(postgres.Open(getPostgresURL(configYAML)), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return db
}
