package viewModels

import (
	"time"
)

// Message - message sent to client
type Message struct {
	ConversationID string    `json:"conversationID,omitempty"`
	AuthorID       string    `json:"authorID,omitempty"`
	Body           string    `json:"body,omitempty"`
	Append         bool      `json:"append,omitempty"`
	TimeStamp      time.Time `json:"timeStamp,omitempty"`
}
