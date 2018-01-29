package models

import "time"

// Message - Message
type Message struct {
	Content   string    `json:"content,omitempty"`
	From      string    `json:"from,omitempty"`
	TimeStamp time.Time `json:"timeStamp,omitempty"`
}

// Conversation - Conversation
type Conversation struct {
	ID           string    `json:"id,omitempty"`
	Participants []string  `json:"participants,omitempty"`
	Name         string    `json:"name,omitempty"`
	Messages     []Message `json:"messages,omitempty"`
}
