package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// InitDB creates a connection to the PostgreSQL database
func InitDB() *sql.DB {
	// Get database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Build connection string
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("✅ Successfully connected to PostgreSQL database!")

	// Create tables if they don't exist
	err = CreateTables(db)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}

	return db
}

// CreateTables creates all necessary database tables
func CreateTables(db *sql.DB) error {
	// SQL to create users table
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	// SQL to create conversations table
	conversationsTable := `
	CREATE TABLE IF NOT EXISTS conversations (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	// SQL to create conversation_members table
	conversationMembersTable := `
	CREATE TABLE IF NOT EXISTS conversation_members (
		conversation_id INTEGER NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		PRIMARY KEY (conversation_id, user_id)
	);
	`

	// SQL to create messages table
	messagesTable := `
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		conversation_id INTEGER NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
		sender_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	// Execute all table creation statements
	statements := []string{usersTable, conversationsTable, conversationMembersTable, messagesTable}

	for _, statement := range statements {
		_, err := db.Exec(statement)
		if err != nil {
			return err
		}
	}

	log.Println("✅ Database tables created successfully!")
	return nil
}