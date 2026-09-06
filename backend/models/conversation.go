package models

import "time"

// Conversation represents a chat conversation between users
type Conversation struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// ConversationMember links a user to a conversation
type ConversationMember struct {
	ConversationID int `json:"conversation_id"`
	UserID         int `json:"user_id"`
}