package db

import "goChat/Server/models"

// IConversationRepository - interface for conversation repository
type IConversationRepository interface {
	GetConversationByID(id string) (models.Conversation, error)
	GetConversationsForUserID(userID string) ([]models.Conversation, error)
	AddConversation(conversation models.Conversation) (string, error)
	AddMessage(message models.Message, convID string) error
}
