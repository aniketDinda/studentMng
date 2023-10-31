package main

import (
	"log"
	"os"
	"studentMng/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8085"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	router.GET("/api-health", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Api is Up"})
	})

	router.Run(":" + port)
}
