package config

import (
	"log"
	"os"
	"time"

	"github.com/JairDavid/Probien-Backend/config/migrations/models"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

func ConnectDB() {
	env := godotenv.Load("vars.env")
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URI_DEV")), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	if err != nil || env != nil {
		log.Fatal(err)
		log.Fatal(env)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(100)
	Database = db
}

//cron job for update pawn orders
func StartCronJobs() {
	//init cron with time zone server
	job := gocron.NewScheduler(time.UTC)

	job.Every(1).Day().Do(func() {
		Database.Exec("CALL update_orders()")
		log.Print("calling stored procedure for update orders...")
	})

	//running job async
	job.StartAsync()
}

//to migrate the models, add this function on main.go before setup all routes
func Migrate() {
	Database.AutoMigrate(
		&models.Category{},
		&models.Customer{},
		&models.BranchOffice{},
		&models.Employee{},
		&models.Product{},
		&models.Status{},
		&models.PawnOrder{},
		&models.Endorsement{},
		&models.SessionLog{},
		&models.ModerationLog{},
		&models.Profile{})
}
