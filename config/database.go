package config

import (
	"github.com/JairDavid/Probien-Backend/config/migrations/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

func ConnectDB() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/probien"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	Database = db
}

func GetDBInstance() *gorm.DB {
	return Database
}

func Migrate() {
	GetDBInstance().AutoMigrate(&models.Category{}, &models.Customer{}, &models.Employee{}, &models.Product{}, &models.Endorsement{}, &models.PawnOrder{}, &models.Status{})
}
