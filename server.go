package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.Run(":9000")
}
