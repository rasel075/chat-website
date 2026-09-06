package handlers

import (
	"chat-website/models"
	"chat-website/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRequest is the request body for user registration
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterResponse is the response after registration
type RegisterResponse struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

// Register handles user registration
func Register(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest

		// Parse request body
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Println("Error parsing request:", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Validate input
		if req.Username == "" || req.Email == "" || req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username, email, and password are required",
			})
			return
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			log.Println("Error hashing password:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error processing password",
			})
			return
		}

		// Insert user into database
		var userID int
		err = db.QueryRow(
			"INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id",
			req.Username, req.Email, hashedPassword,
		).Scan(&userID)

		if err != nil {
			log.Println("Error inserting user:", err)
			// Check if it's a duplicate username or email
			c.JSON(http.StatusConflict, gin.H{
				"error": "Username or email already exists",
			})
			return
		}

		// Return success response
		user := models.User{
			ID:       userID,
			Username: req.Username,
			Email:    req.Email,
		}

		c.JSON(http.StatusCreated, RegisterResponse{
			Message: "User registered successfully",
			User:    user,
		})
	}
}

/*
Part	Explanation
RegisterRequest	Struct that maps JSON request body to Go struct
binding:"required"	Ensures these fields are not empty
c.ShouldBindJSON()	Parses JSON from request
utils.HashPassword()	Converts password to secure hash
db.QueryRow()	Executes SQL and gets one result
RETURNING id	SQL syntax to return the new user's ID
http.StatusCreated	HTTP 201 status (resource created)
http.StatusConflict	HTTP 409 status (duplicate exists)
*/
