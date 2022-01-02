package main

import (
	"github.com/JairDavid/Probien-Backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// config.ConnectDB()
	// config.Migrate()

	utils.Setup(server)
	server.Run(":9000")
}
