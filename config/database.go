package config

import (
	"log"
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
			db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=root dbname=probien port=5432 sslmode=disable"), &gorm.Config{})
			database = db
			if err != nil {
				log.Fatal(err)
			} else {
				log.Printf("the database has been connected successfuly")
			}
		})
}

func GetDBInstance() *gorm.DB {
	return database
}

func Migrate() {
	GetDBInstance().AutoMigrate(&models.Category{}, &models.Customer{}, &models.Employee{}, &models.Product{}, &models.Endorsement{}, &models.PawnOrder{}, &models.Status{})
}
