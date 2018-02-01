package services

import (
	"goChat/Server/db"
	"goChat/Server/models"
	"goChat/Server/utils"

	"goChat/Server/viewModels"
	"log"
	"net/http"
)

// ConversationService - Conversation
type ConversationService struct {
	conversationRepo db.IConversationRepository
	userService      *UserService
}

// NewConversationService - Creates new instances of conversation service
func NewConversationService(repo db.IConversationRepository, userSvc *UserService) *ConversationService {
	return &ConversationService{
		conversationRepo: repo,
		userService:      userSvc,
	}
}

// GetConversationByID - gets the conversation by id
func (service *ConversationService) GetConversationByID(id string) (models.Conversation, error) {
	return service.conversationRepo.GetConversationByID(id)
}

// GetConversationHandler - api handler function for get conversation for user
func (service *ConversationService) GetConversationHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(utils.RequestContextKeyUser).(models.User)
	convList, err := service.conversationRepo.GetConversationsForUserID(user.ID)
	if err != nil {
		utils.JSONInternalServerErrorResponse(w, err)
		return
	}
	var convModelList []viewModels.Conversation
	for _, c := range convList {
		var pInfoList []viewModels.ParticipantInfo
		for _, p := range c.Participants {
			pUser, err := service.userService.GetUserByID(p)
			if err != nil {
				log.Fatalln(err)
				continue
			}
			pInfo := viewModels.ParticipantInfo{
				Email:         pUser.Email,
				FirstName:     pUser.FirstName,
				LastName:      pUser.LastName,
				ParticipantID: pUser.ID,
			}

			pInfoList = append(pInfoList, pInfo)
		}

		convItem := viewModels.Conversation{
			ConversationID: c.ID,
			Participants:   pInfoList,
			Name:           c.Name,
		}

		convModelList = append(convModelList, convItem)
	}

	utils.JSONSuccessResponse(w, convModelList)
}
