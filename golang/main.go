package main

import (
	// "app/controllers"

	"codechalllenge/database"
	"fmt"
	"log"

	// "models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	seedDatabase()
	serveApplication()
}

func seedDatabase() {
	// TODO:
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer db.Close()
}

func serveApplication() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":3001")
	fmt.Println("Server running on port 3000")
}
