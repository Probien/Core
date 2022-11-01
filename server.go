package main

import (
	"flag"
	"log"

	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	migrate := flag.Bool("migrate", false, "migrate datamodel structs and stored procedures to database")
	flag.Parse()
	server := gin.Default()
	config.ConnectDB()

	if *migrate {
		config.Migrate()
		log.Print("migrated all models")
	}

	config.StartCronJobs()
	config.ConnectRedis()

	router.Setup(server)

	if err := server.Run(":9000"); err != nil {
		log.Print(err.Error())
	}

}
