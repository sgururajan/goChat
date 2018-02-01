package utils

import (
	"goChat/Server/models"
	"goChat/Server/viewModels"
)

// ConvertToViewModelMessage - Convert the db message model to view model messages
func ConvertToViewModelMessage(dbMsg models.Message) viewModels.Message {
	return viewModels.Message{
		ConversationID: dbMsg.ConversationID,
		AuthorID:       dbMsg.Author,
		Body:           dbMsg.Body,
		TimeStamp:      dbMsg.TimeStamp,
		Append:         true,
	}
}
