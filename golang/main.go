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
	//seedDatabase()
	serveApplication()
}

func seedDatabase() {
	// TODO: Instruction 2
	fmt.Println("seedDatabase")

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	fmt.Println("Connect")
	err = database.InsertExample(db, 50777, "2025-03-05 12:00:00+00")
	if err != nil {
		log.Fatal("erro ao escrever:", err)
	}

	fmt.Println("InsertExample")

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
