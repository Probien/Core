package config

import (
	"time"

	"github.com/JairDavid/Probien-Backend/config/migrations/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/probien"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	configs, err := db.DB()
	if err != nil {
		panic(err)
	}
	configs.SetMaxIdleConns(100)
	configs.SetMaxOpenConns(50)
	configs.SetConnMaxLifetime(time.Minute * 5)

	return db
}

func Migrate() {
	conn := ConnectDB()
	conn.AutoMigrate(&models.Category{}, &models.Customer{}, &models.Employee{}, &models.Product{}, &models.Endorsement{}, &models.PawnOrder{}, &models.Status{})
}
