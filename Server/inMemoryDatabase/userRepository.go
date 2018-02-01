package inMemoryDatabase

import (
	"fmt"
	"goChat/Server/models"
	"goChat/Server/utils"
	"log"
)

// UserRepository - UserRepository
type UserRepository struct {
	users map[string]models.User
}

// NewUserRepository - NewUserRepository
func NewUserRepository() *UserRepository {
	log.Println("Intializing in memory user repository")
	return &UserRepository{
		users: make(map[string]models.User),
	}
}

func (repo *UserRepository) isUserExists(user models.User) bool {
	for _, value := range repo.users {
		if value.Email == user.Email {
			return true
		}
	}

	return false
}

// Create - Creates New User
func (repo *UserRepository) Create(user models.User) (string, error) {
	if repo.isUserExists(user) {
		return "", fmt.Errorf("user with email %s already exists", user.Email)
	}

	user.ID, _ = utils.GenerateNewGUID()
	repo.users[user.ID] = user
	return "", nil
}

// GetUserByEmail - GetUserByEmail
func (repo *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	found := false
	for _, val := range repo.users {
		if val.Email == email {
			user = val
			found = true
			break
		}
	}

	if !found {
		return user, fmt.Errorf("user with email %s does not exists", email)
	}

	return user, nil

}

// GetUserByID - GetUserByID
func (repo *UserRepository) GetUserByID(id string) (models.User, error) {
	user, ok := repo.users[id]
	if !ok {
		return user, fmt.Errorf("user record with ID %s does not exists", id)
	}
	return user, nil
}
