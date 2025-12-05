package main

import (
	"log"
	"magic-link-auth/internal/auth"
	"magic-link-auth/internal/config"
	"magic-link-auth/internal/database"
	"magic-link-auth/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectMySQL()

	database.DB.AutoMigrate(&models.User{}, &models.MagicToken{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context){
		c. JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth.RegisterAuthRoutes(r)

	log.Println("Server running on : 8080")
	r.Run()
}