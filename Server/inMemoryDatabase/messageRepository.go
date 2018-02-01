package inMemoryDatabase

import (
	"goChat/Server/models"
	"goChat/Server/utils"
	"log"

	"time"
)

// MessageRepository - MessageRepository
type MessageRepository struct {
	messages map[string]models.Message
}

// NewMessageRepository - creates new instance of MessageRepository
func NewMessageRepository() *MessageRepository {
	log.Println("Intializing in memory message repository")
	return &MessageRepository{
		messages: make(map[string]models.Message),
	}
}

// GetMessagesByConversation - GetMessagesByConversation
func (repo *MessageRepository) GetMessagesByConversation(conversationID string, page, count int) ([]models.Message, error) {
	var messages []models.Message
	pageStart := (page - 1) * count
	counter := 0
	for _, m := range repo.messages {
		if m.ConversationID == conversationID {
			if counter >= pageStart && len(messages) < count {
				messages = append(messages, m)
			}
			counter = counter + 1
		}
	}

	return messages, nil
}

// AddMessage - Add message
func (repo *MessageRepository) AddMessage(message models.Message) error {
	// message.MessageID = NewObjectID()
	// message.TimeStamp = time.Now()
	// err := repo.collection.Insert(message)
	// return err

	message.MessageID, _ = utils.GenerateNewGUID()
	message.TimeStamp = time.Now()
	repo.messages[message.MessageID] = message

	return nil
}

// UpdateMessageAsRead - UpdateMessageAsRead
func (repo *MessageRepository) UpdateMessageAsRead(msgID, participantID string) {
	msg, ok := repo.messages[msgID]
	if !ok {
		return
	}

	for _, p := range msg.ParticipantsState {
		if p.ParticipantID == participantID {
			p.IsRead = true
		}
	}
}
