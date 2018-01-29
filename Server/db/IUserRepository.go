package db

import (
	"goChat/Server/models"
)



// IUserRespository - interface for user repository functions
type IUserRespository interface {
	Create(user models.User) (string, error)
	GetUserByEmail(email string) (models.User, error)
	// Login(userName, password string) (string, error)
	// Logout(userID, sessionID string)
	// AddInvitationByUserId(userID string) error
	// AddInvitationByUser(user models.User) error
	// AcceptInvitation(acceptedUserID string) error
	// ClientHealthCheck(userID, sessionID string)
}
