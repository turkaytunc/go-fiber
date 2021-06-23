package database

import (
	"github.com/turkaytunc/go-web-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=pass123 dbname=test port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("cannot connect db")
	}

	db.AutoMigrate(&models.User{})

	DB = db
	return db
}
