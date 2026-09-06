package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

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