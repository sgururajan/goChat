package db

import "goChat/Server/models"

// IMessageRepository - interface for Message repository
type IMessageRepository interface {
	GetMessagesByConversation(conversationID string, page, count int) ([]models.Message, error)
	AddMessage(message models.Message) error
	UpdateMessageAsRead(msgID, participantID string)
}
