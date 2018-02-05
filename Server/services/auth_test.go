package services

import (
	"bytes"
	"goChat/Server/inMemoryDatabase"
	"goChat/Server/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewAuthService(t *testing.T) {
	repo := getInMemoryUserRepo()
	authService := NewAuthService(repo)
	//userId, _ := createTestUser(repo)
	//user, _ := repo.GetUserByID(userId)
	if authService == nil {
		t.Errorf("failed to initialize auth service")
	}
}

func TestAuthService_AuthenticateHandler(t *testing.T) {
	repo := getInMemoryUserRepo()
	authService := NewAuthService(repo)
	createTestUser(repo)
	// user, _ := repo.GetUserByID(userId)
	jsonCredential := []byte(`{"userName":"siva@gochat.com","password":"hello"}`)
	//data := url.Values{}
	//data.Add("userName", user.Email)
	//data.Add("password", "hello")
	req, err := http.NewRequest("GET", "/login", bytes.NewBuffer(jsonCredential))
	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(authService.AuthenticateHandler)

	handler.ServeHTTP(resRec, req)

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("handler returned unexpected result")
	}

	jsonCredential = []byte(`{"userName":"siva1@gochat.com","password":"hello"}`)
	req, err = http.NewRequest("GET", "/login", bytes.NewBuffer(jsonCredential))
	resRec = httptest.NewRecorder()
	handler = http.HandlerFunc(authService.AuthenticateHandler)
	handler.ServeHTTP(resRec, req)
	if status := resRec.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned unexpected result for nagative user")
	}
}

func createTestUser(repo *inMemoryDatabase.UserRepository) (string, error) {
	user := models.User{
		Email:          "siva@gochat.com",
		FirstName:      "siva",
		LastName:       "siva",
		PasswordHashed: "hello",
		NickName:       "siva",
	}

	return repo.Create(user)
}

func getInMemoryUserRepo() *inMemoryDatabase.UserRepository {
	return inMemoryDatabase.NewUserRepository()
}
