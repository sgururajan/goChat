package models

import (
	"time"
)

// UserOnlineStatus - users online status
type UserOnlineStatus uint

const (
	// Online - user is online and connected to server
	Online UserOnlineStatus = iota
	// Idle - user is online but detected no activity for a while
	Idle
	// Offline - user logged out or connection broken
	Offline
)

// User - User Model
type User struct {
	ID             string           `json:"id, omitempty"`
	FirstName      string           `json:"firstName, omitempty"`
	LastName       string           `json:"lastName, omitempty"`
	Email          string           `json:"email, omitempty"`
	NickName       string           `json:"nickName, omitempty"`
	PasswordHashed string           `json:"passwordHashed, omitempty"`
	CreatedOn      time.Time        `json:"createdOn, omitempty"`
	Status         UserOnlineStatus `json:"status, omitempty"`
	Contacts       []UserContact    `json:"contacts, omitempty"`
}

// UserContact  - list of contacts for each user
type UserContact struct {
	UserID           string    `json:"userID, omitempty"`
	InviteAccepted   bool      `json:"inviteAccepted, omitempty"`
	InviteAcceptedOn time.Time `json:"inviteAcceptedOn, omitempty"`
}

// UserCredential - user credentials for authentication
type UserCredential struct {
	UserName string
	Password string
}
