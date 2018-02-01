package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"goChat/Server/db"
	"goChat/Server/models"
	"goChat/Server/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const signingKey = "APIKeyWHICHISNotTOOSECRET"

var keyGetter = func(t *jwt.Token) (interface{}, error) {
	return []byte(signingKey), nil
}

// TokenClaims - TokenClaims
type TokenClaims struct {
	UserName string `json:"userName"`
	jwt.StandardClaims
}

// AuthService - AuthService
type AuthService struct {
	userRepository db.IUserRepository
}

// Token - Token
type Token struct {
	AccessToken string `json:"access_token"`
}

// LoginResponse - LoginResponse
type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NickName    string `json:"nickName"`
}

// NewAuthService - create new instance of auth service
func NewAuthService(userRepo db.IUserRepository) *AuthService {
	authService := AuthService{
		userRepository: userRepo,
	}

	return &authService
}

// AuthenticateHandler - handler function for authentication
func (auth *AuthService) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	var userCredential models.UserCredential
	_ = json.NewDecoder(r.Body).Decode(&userCredential)
	user, err := auth.userRepository.GetUserByEmail(userCredential.UserName)
	if err != nil {
		msg := fmt.Sprintf("The provided user name: %s cannot be found", userCredential.UserName)
		utils.JSONUnAuthorizedResponse(w, msg)
		log.Printf(msg)
		return
	}

	pwdMatch := utils.CheckPasswordHash(userCredential.Password, user.PasswordHashed)
	if !pwdMatch {
		log.Printf("The provided password does not match")
		utils.JSONUnAuthorizedResponse(w, "Invalid credentials")
		return
	}

	token, err := generateToken(user)

	if err != nil {
		utils.JSONInternalServerErrorResponse(w, fmt.Sprintf("Error while creating token: %s", err))
		log.Printf("Error while creating token: %s", err)
		return
	}

	loginResponse := LoginResponse{
		AccessToken: token,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		NickName:    user.NickName,
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error while converting token to json object")
		return
	}

	utils.JSONSuccessResponse(w, loginResponse)
}

// AuthenticationMiddleware - middleware funciton that authenticates the request using JWT
func (auth *AuthService) AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.validateRequest(r)
		if err != nil {
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)
			return
		}
		user, err := auth.parseUserFromRequest(token)
		if err != nil {
			http.Error(w, "Unable to get user information", http.StatusInternalServerError)
			return
		}
		rWithContext := r.WithContext(context.WithValue(r.Context(), utils.RequestContextKeyUser, user))
		*r = *rWithContext
		next(w, r)
	})
}

func (auth *AuthService) validateRequest(r *http.Request) (*jwt.Token, error) {
	tokenString, err := extractTokenFromHeader(r)

	if err != nil {
		return &jwt.Token{}, fmt.Errorf("Error while extracting auth token: %s", err)
	}

	if tokenString == "" {
		return &jwt.Token{}, fmt.Errorf("No authentication token found")
	}

	token, err := jwt.Parse(tokenString, keyGetter)

	if err != nil {
		return &jwt.Token{}, fmt.Errorf("Error while validating auth token")
	}

	if !token.Valid {
		return &jwt.Token{}, fmt.Errorf("Invalid token")
	}

	return token, nil
}

func (auth *AuthService) parseUserFromRequest(token *jwt.Token) (models.User, error) {
	user := models.User{}

	claims, ok := token.Claims.(*TokenClaims)

	if !ok {
		return user, fmt.Errorf("Unable to process claims")
	}

	user, err := auth.userRepository.GetUserByEmail(claims.UserName)

	if err != nil {
		return user, fmt.Errorf("Unable to find user given in token. Error: %s", err)
	}

	return user, nil
}

func extractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", nil
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 && strings.ToUpper(headerParts[0]) != "BEARER" {
		return "", errors.New("Authorization header not in correct format")
	}

	return headerParts[1], nil
}

func generateToken(user models.User) (string, error) {

	claims := TokenClaims{
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(2)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signingKey))

}
