package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// config.ConnectDB()
	//config.Migrate()

	//utils.Setup(server)
	server.Run(":9000")
}
