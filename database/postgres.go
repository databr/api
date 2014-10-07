package database

import (
	"log"
	"os"

	"github.com/databr/api/models"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

var Postgres gorm.DB

func NewPostgres() gorm.DB {
	var err error

	databaseUrl, _ := pq.ParseURL(os.Getenv("DATABASE_URL"))
	Postgres, err = gorm.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatalln(err.Error())
	}

	Postgres.LogMode(true)
	Postgres.AutoMigrate(models.App{})
	Postgres.AutoMigrate(models.User{})
	return Postgres
}
