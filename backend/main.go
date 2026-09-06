package main

import (
	"chat-website/database"
	"chat-website/routes"
	"chat-website/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables from .env file
	utils.LoadEnv()

	// Initialize database connection
	db := database.InitDB()
	defer db.Close()

	// Create a new Gin router
	router := gin.Default()

	// Setup all routes
	routes.SetupRoutes(router, db)

	// Define a simple GET endpoint for testing
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Chat website backend is running!",
		})
	})

	// Define a simple health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}