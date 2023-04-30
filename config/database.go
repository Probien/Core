package config

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/JairDavid/Probien-Backend/config/migration/model"
	"github.com/go-co-op/gocron"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectDB() {

	db, err := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/probien?sslmode=disable"), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(100)

	database = db
}

func GetConnection() *gorm.DB {
	return database
}

// StartCronJobs cron job for update pawn orders
func StartCronJobs() {
	//init cron with time zone server
	job := gocron.NewScheduler(time.UTC)

	_, cronErr := job.Every(1).Day().Do(func() {
		database.Exec("CALL update_orders()")
		log.Print("calling stored procedure for update orders...")
	})

	if cronErr != nil {
		panic(cronErr.Error())
	}

	//running job async
	job.StartAsync()
}

// Migrate to migrate the models and stored procedures, add flag -migrate=true
func Migrate() {
	sp1, sp1Err := ioutil.ReadFile("./config/migration/stored procedures/sessions.sql")
	sp2, sp2Err := ioutil.ReadFile("./config/migration/stored procedures/moderation.sql")
	sp3, sp3Err := ioutil.ReadFile("./config/migration/stored procedures/order_dates.sql")

	if sp1Err != nil || sp2Err != nil || sp3Err != nil {
		panic(sp1Err.Error() + sp2Err.Error() + sp3Err.Error())
	}

	err := database.AutoMigrate(
		&model.Category{},
		&model.Customer{},
		&model.BranchOffice{},
		&model.Employee{},
		&model.Role{},
		&model.Product{},
		&model.Status{},
		&model.PawnOrder{},
		&model.Endorsement{},
		&model.SessionLog{},
		&model.ModerationLog{},
		&model.Profile{},
	)
	if err != nil {
		panic(err.Error())
	}

	database.Exec(string(sp1))
	database.Exec(string(sp2))
	database.Exec(string(sp3))
}
