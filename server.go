package main

import (
	"log"
	"time"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/router"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

func main() {
	server := gin.Default()
	config.ConnectDB()
	//config.Migrate() for migrate all models

	router.Setup(server)
	job := gocron.NewScheduler(time.UTC)

	job.Every(1).Day().Do(func() {
		config.Database.Exec("CALL update_orders()")
		log.Print("calling stored procedure for update orders")
	})

	//running job async
	job.StartAsync()

	server.Run(":9000")
}
