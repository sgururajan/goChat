package services

import (
	"goChat/Server/db"
	"goChat/Server/models"
)

// Conversation - Conversation
type ConversationService struct {
	conversationRepo db.IConversationRepository
}

// NewConversation - Creates new instances of conversation service
func NewConversationService(repo db.IConversationRepository) *ConversationService {
	return &ConversationService{
		conversationRepo: repo,
	}
}

// GetConversationByID - gets the conversation by id
func (conv *ConversationService) GetConversationByID(id string) (models.Conversation, error) {
	return conv.conversationRepo.GetConversationByID(id)
}

// GetConversationForUserID - gets list of conversations for user Id
func (conv *ConversationService) GetConversationForUserID(userId string) ([]models.Conversation, error) {
	return conv.conversationRepo.GetConversationsForUserID(userId)
}
