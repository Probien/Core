package main

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/JairDavid/Probien-Backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	config.ConnectDB()
	//config.Migrate() for migrate all models
	config.StartCronJobs()
	config.ConnectRedis()

	router.Setup(server)
	server.Run(":9000")
}
