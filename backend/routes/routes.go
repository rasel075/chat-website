package routes

import (
	"chat-website/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// SetupRoutes defines all API routes
func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Authentication routes
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			// POST /api/auth/register - Register a new user
			auth.POST("/register", handlers.Register(db))
		}
	}
}
/*
Explanation:

router.Group() groups routes under a prefix (/api)
auth.POST() creates a POST endpoint
handlers.Register(db) passes the database to the handler
*/