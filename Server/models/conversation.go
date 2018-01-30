package models

// Conversation - Conversation
type Conversation struct {
	ID           string   `json:"id,omitempty" bson:"_id"`
	Participants []string `json:"participants,omitempty"`
	Name         string   `json:"name,omitempty"`
}
