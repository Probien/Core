package config

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/JairDavid/Probien-Backend/config/migration/model"
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
		log.Fatal(err.Error() + env.Error())
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

	_, _ = job.Every(1).Day().Do(func() {
		Database.Exec("CALL update_orders()")
		log.Print("calling stored procedure for update orders...")
	})

	//running job async
	job.StartAsync()
}

//to migrate the models and stored procedures, add flag -migrate=true
func Migrate() {
	sp1, sp1Err := ioutil.ReadFile("./config/migration/stored procedures/sessions.sql")
	sp2, sp2Err := ioutil.ReadFile("./config/migration/stored procedures/moderation.sql")
	sp3, sp3Err := ioutil.ReadFile("./config/migration/stored procedures/order_dates.sql")

	if sp1Err != nil || sp2Err != nil || sp3Err != nil {
		panic(sp1Err.Error() + sp2Err.Error() + sp3Err.Error())
	}

	err := Database.AutoMigrate(
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
		log.Print(err)
	}

	Database.Exec(string(sp1))
	Database.Exec(string(sp2))
	Database.Exec(string(sp3))
}
