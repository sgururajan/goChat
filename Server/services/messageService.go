package services

import (
	"goChat/Server/db"
	"goChat/Server/models"
)

// MessageService - MessageService
type MessageService struct {
	repo db.IMessageRepository
}

// NewMessageService - Creates new instance of MessageService
func NewMessageService(repository db.IMessageRepository) *MessageService {
	return &MessageService{
		repo: repository,
	}
}

// AddMessage - AddMessage
func (svc *MessageService) AddMessage(message models.Message) error {
	return svc.repo.AddMessage(message)
}
