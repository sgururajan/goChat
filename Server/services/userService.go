package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goChat/Server/db"
	"goChat/Server/models"
	"goChat/Server/utils"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// UserService - UserService
type UserService struct {
	userRepo db.IUserRepository
}

// NewUserService - creates new instance of user service
func NewUserService(userRepository db.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepository,
	}
}

// SignupHandler - Signup Handler
func (service *UserService) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := service.createUserInternal(user)
	if err != nil {
		utils.FailOnServerError(w, fmt.Sprintf("%s", err))
	}
}

// SignupHandlerWithNext - SignupHandlerWithNext
func (service *UserService) SignupHandlerWithNext(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		err := service.createUserInternal(user)
		if err != nil {
			utils.FailOnServerError(w, fmt.Sprintf("%s", err))
		}

		credentials := models.UserCredential{
			UserName: user.Email,
			Password: user.PasswordHashed,
		}
		log.Printf("Credentials: %s", credentials)
		credBytes, err := json.Marshal(credentials)
		if err != nil {
			utils.FailOnServerError(w, fmt.Sprintf("User created but unable to signin the user: %s", err))
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(credBytes))
		next(w, r)
	})
}

// GetContactListHandler - GetContactListHandler
func (service *UserService) GetContactListHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(utils.RequestContextKeyUser).(models.User)
	utils.JSONSuccessResponse(w, user.Contacts)
}

func (service *UserService) createUserInternal(user models.User) error {
	dbUser, err := service.userRepo.GetUserByEmail(user.Email)
	if err == nil && dbUser.Email == user.Email {
		return fmt.Errorf("The user with email %s already exists", user.Email)
	}

	pwdHashed, err := utils.HashPassword(user.PasswordHashed)
	if err != nil {
		return fmt.Errorf("Error while processing request: %s", err)
	}

	user.PasswordHashed = pwdHashed
	user.CreatedOn = time.Now()
	_, err = service.userRepo.Create(user)

	if err != nil {
		return fmt.Errorf("Error while creating user: %s", err)
	}

	return nil
}

// GetUserByID - GetUserByID
func (service *UserService) GetUserByID(id string) (models.User, error) {
	return service.GetUserByID(id)
}
