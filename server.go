package main

import (
	"flag"
	"github.com/go-co-op/gocron"
	"log"
	"time"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	migrate := flag.Bool("migrate", false, "migrate struct models and stored procedures to database")
	flag.Parse()
	pgClient := config.NewPostgresConnection("postgres://postgres:root@localhost:5432/probien?sslmode=disable")

	if *migrate {
		pgClient.Migrate()
	}

	scheduler := config.NewScheduler(gocron.NewScheduler(time.Local), pgClient)
	scheduler.StartCronJobs()

	// NewRedisClient receive host and password
	config.NewRedisClient("localhost:6379", "")

	server := gin.Default()
	router.Setup(server)

	if err := server.Run(":9000"); err != nil {
		log.Fatalln(err)
	}

}
