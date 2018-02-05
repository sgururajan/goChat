package inMemoryDatabase

import (
	"goChat/Server/models"
	"testing"
)

func TestNewUserRepository(t *testing.T) {
	t.Parallel()
	repo := getTestUserRepository()
	if repo.users == nil {
		t.Errorf("user repository map not initialized properly")
	}
}

func TestUserRepository_Create(t *testing.T) {
	t.Parallel()
	repo := getTestUserRepository()
	user := models.User{
		Email:          "siva@gochat.com",
		FirstName:      "siva",
		LastName:       "siva",
		PasswordHashed: "hello",
		NickName:       "siva",
	}

	got, err := repo.Create(user)
	if len(got) == 0 || err != nil {
		t.Errorf("unable to create user: %s", err)
	}

	// test if the password is hashed
	user, _ = repo.GetUserByID(got)
	if user.PasswordHashed == "hello" {
		t.Errorf("password is not hashed properly")
	}

	// try creating with same user. it should error out now
	got, err = repo.Create(user)
	if err == nil {
		t.Errorf("creating with same user email failed")
	}

	if len(repo.users) > 1 {
		t.Errorf("the repository should have one user at this time")
	}

}

func TestUserRepository_GetUserByID(t *testing.T) {
	t.Parallel()
	repo := getTestUserRepository()
	userId, _ := createTestUser(repo)

	user, err := repo.GetUserByID(userId)
	if err != nil {
		t.Errorf("failed to find user by id")
	}

	if user.Email != "siva@gochat.com" {
		t.Errorf("the user object found is not expected")
	}
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	t.Parallel()
	repo := getTestUserRepository()
	createTestUser(repo)

	user, err := repo.GetUserByEmail("siva@gochat.com")
	if err != nil {
		t.Errorf("failed to find the created user")
	}

	if user.Email != "siva@gochat.com" {
		t.Errorf("the user object fould is not expected")
	}

}

func createTestUser(repo *UserRepository) (string, error) {
	user := models.User{
		Email:          "siva@gochat.com",
		FirstName:      "siva",
		LastName:       "siva",
		PasswordHashed: "hello",
		NickName:       "siva",
	}

	return repo.Create(user)
}

func getTestUserRepository() *UserRepository {
	return NewUserRepository()
}
