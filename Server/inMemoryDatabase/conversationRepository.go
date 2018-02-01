package inMemoryDatabase

import (
	"fmt"
	"goChat/Server/models"
	"goChat/Server/utils"
	"log"
	"sort"
)

// ConversationRepository - ConversationRepository
type ConversationRepository struct {
	conversations map[string]models.Conversation
}

// NewConversationRepository - Creates new instance of ConversationRepository
func NewConversationRepository() *ConversationRepository {
	log.Println("Intializing in memory conversation repository")
	return &ConversationRepository{
		conversations: make(map[string]models.Conversation),
	}
}

// GetConversationsForUserID -  GetConversationsForUserID
func (repo *ConversationRepository) GetConversationsForUserID(userID string) ([]models.Conversation, error) {
	var conversation []models.Conversation

	for _, c := range repo.conversations {
		i := sort.SearchStrings(c.Participants, userID)
		if i < len(c.Participants) && c.Participants[i] == userID {
			conversation = append(conversation, c)
		}
	}

	return conversation, nil
}

// AddConversation - AddConversation
func (repo *ConversationRepository) AddConversation(conversation models.Conversation) (string, error) {
	cID, _ := utils.GenerateNewGUID()
	conversation.ID = cID
	repo.conversations[cID] = conversation
	return "", nil
}

// AddMessage - AddMessage
func (repo *ConversationRepository) AddMessage(message models.Message, convID string) error {

	return nil
}

// GetConversationByID - GetConversationByID
func (repo *ConversationRepository) GetConversationByID(id string) (models.Conversation, error) {
	conv, ok := repo.conversations[id]
	if !ok {
		return models.Conversation{}, fmt.Errorf("Conversation not found")
	}
	return conv, nil
}
