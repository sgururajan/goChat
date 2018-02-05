package services

import (
	"goChat/Server/db"
	"goChat/Server/models"
	"goChat/Server/utils"
	"goChat/Server/viewModels"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// GetMessageHandler - Gets the message for conversation Id
func (svc *MessageService) GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	convID := args["conversationId"]
	skip, err := strconv.Atoi(args["skip"])
	if err != nil {
		skip = 0
	}
	count, err := strconv.Atoi(args["count"])
	if err != nil {
		count = 10
	}
	msgList, err := svc.repo.GetMessagesByConversation(convID, skip, count)
	if err != nil {
		utils.JSONInternalServerErrorResponse(w, err)
		return
	}

	var resultMsg []viewModels.Message
	for _, m := range msgList {
		resultMsg = append(resultMsg, utils.ConvertToViewModelMessage(m))
	}

	utils.JSONSuccessResponse(w, resultMsg)
}

// UpdateMessageAsRead - UpdateMessageAsRead
func (svc *MessageService) UpdateMessageAsRead(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	user := r.Context().Value(utils.RequestContextKeyUser).(models.User)
	msgId := args["messageId"]
	go svc.repo.UpdateMessageAsRead(msgId, user.ID)
	utils.JSONSuccessResponse(w, nil)
}
