package main

import (
	"log"
	"magic-link-auth/internal/config"
	"magic-link-auth/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectMySQL()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context){
		c. JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server running on : 8080")
	r.Run()
}