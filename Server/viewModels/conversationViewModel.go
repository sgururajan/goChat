package viewModels

// ParticipantInfo - Participant Info
type ParticipantInfo struct {
	FirstName     string `json:"firstName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
	Email         string `json:"email,omitempty"`
	ParticipantID string `json:"participantID,omitempty"`
}

// Conversation - Conversation model sent to/from client
type Conversation struct {
	ConversationID string            `json:"conversationID,omitempty"`
	Participants   []ParticipantInfo `json:"participants,omitempty"`
	Name           string            `json:"name,omitempty"`
}
