package postgres

import (
	"github.com/JairDavid/Probien-Backend/internal/infra/resource/postgres/migration/model"
	"gorm.io/driver/postgres"
	"log"
	"os"

	"gorm.io/gorm"
)

type Client struct {
	Conn *gorm.DB
}

// NewPostgresConnection receive a formatted string "postgres://postgres:user@ip:port/database_name?sslmode=disable"
func NewPostgresConnection(dsn string) *Client {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(100)

	return &Client{
		Conn: db,
	}
}

func (p *Client) GetConnection() *gorm.DB {
	return p.Conn
}

// Migrate to migrate the models and stored procedures, add flag -migrate=true
func (p *Client) Migrate() {

	sessions, err := os.ReadFile("./config/migration/stored procedures/sessions.sql")
	if err != nil {
		log.Fatalln(err)
	}

	moderation, err := os.ReadFile("./config/migration/stored procedures/moderation.sql")
	if err != nil {
		log.Fatalln(err)
	}

	orderDates, err := os.ReadFile("./config/migration/stored procedures/order_dates.sql")
	if err != nil {
		log.Fatalln(err)
	}

	err = p.Conn.AutoMigrate(
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
		log.Fatalln(err)
		return
	}

	p.Conn.Exec(string(sessions))
	p.Conn.Exec(string(moderation))
	p.Conn.Exec(string(orderDates))
	log.Print("migrated all models")
}
