package main

import (
	// "app/controllers"

	"codechalllenge/database"
	"fmt"

	// "models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// loadDatabase()
	// serveApplication()
}

func loadDatabase() {
	database.Connect()
	// database.Database.AutoMigrate(&models.EnergyData{})
}

func serveApplication() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":3000")
	fmt.Println("Server running on port 3000")
}
