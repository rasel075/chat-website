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

// LoginRequest is the request body for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse is the response after login
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User    struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
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

// Login handles user login
func Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest

		// Parse request body
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Println("Error parsing request:", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Validate input
		if req.Email == "" || req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Email and password are required",
			})
			return
		}

		// Query user by email
		var userID int
		var username string
		var passwordHash string

		err := db.QueryRow(
			"SELECT id, username, password_hash FROM users WHERE email = $1",
			req.Email,
		).Scan(&userID, &username, &passwordHash)

		if err == sql.ErrNoRows {
			log.Println("User not found:", req.Email)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}

		if err != nil {
			log.Println("Error querying user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error",
			})
			return
		}

		// Check if password is correct
		if !utils.CheckPassword(passwordHash, req.Password) {
			log.Println("Invalid password for user:", req.Email)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}

		// Generate JWT token
		token, err := utils.GenerateToken(userID, username, req.Email)
		if err != nil {
			log.Println("Error generating token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error generating token",
			})
			return
		}

		// Return success response with token
		c.JSON(http.StatusOK, LoginResponse{
			Message: "Login successful",
			Token:   token,
			User: struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
				Email    string `json:"email"`
			}{
				ID:       userID,
				Username: username,
				Email:    req.Email,
			},
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
