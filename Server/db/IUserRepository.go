package db

import (
	"goChat/Server/models"
)

// IUserRepository - interface for user repository functions
type IUserRepository interface {
	Create(user models.User) (string, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id string) (models.User, error)
	// Login(userName, password string) (string, error)
	// Logout(userID, sessionID string)
	// AddInvitationByUserId(userID string) error
	// AddInvitationByUser(user models.User) error
	// AcceptInvitation(acceptedUserID string) error
	// ClientHealthCheck(userID, sessionID string)
}
