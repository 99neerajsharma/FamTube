package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func getPostgresURL() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		"localhost", "postgres", "postgres", "postgres", "5433")
}
func PostgresInitializer() *gorm.DB {
	db, err := gorm.Open(postgres.Open(getPostgresURL()), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return db
}
