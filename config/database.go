package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

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
			environment := godotenv.Load()
			if environment == nil {
				log.Fatal(environment)
			}
			URI := os.Getenv("DATABASE_URI_DEV")
			db, err := gorm.Open(postgres.Open(URI), &gorm.Config{})
			database = db
			if err != nil {
				log.Fatal(err.Error())
			} else {
				log.Printf("the database has been connected successfuly")
			}

		})
}

func GetDBInstance() *gorm.DB {
	return database
}
