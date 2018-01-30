package models

import "time"

// MessageParicipantState - MessageParicipantState
type MessageParicipantState struct {
	ParticipantID string `json:"participantID,omitempty"`
	IsRead        bool   `json:"isRead,omitempty"`
}

// Message - Message
type Message struct {
	MessageID         string                   `json:"messageID,omitempty" bson:"_id"`
	ConversationID    string                   `json:"conversationID,omitempty"`
	Body              string                   `json:"body,omitempty"`
	Author            string                   `json:"author,omitempty"`
	TimeStamp         time.Time                `json:"timeStamp,omitempty"`
	ParticipantsState []MessageParicipantState `json:"participantsState,omitempty"`
}
