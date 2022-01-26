package config

import (
	"sync"

	"github.com/JairDavid/Probien-Backend/config/migrations/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	database *gorm.DB
	lock     sync.Once
)

func ConnectDB() {
	lock.Do(
		func() {
			db, err := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/probien"), &gorm.Config{})
			database = db
			if err != nil {
				panic(err)
			}
		})
}

func GetDBInstance() *gorm.DB {
	return database
}

func Migrate() {
	GetDBInstance().AutoMigrate(&models.Category{}, &models.Customer{}, &models.Employee{}, &models.Product{}, &models.Endorsement{}, &models.PawnOrder{}, &models.Status{})
}
