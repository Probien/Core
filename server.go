package main

import (
	"github.com/JairDavid/Probien-Backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	config.ConnectDB()
	config.Migrate()
	server.Run(":9000")
}
