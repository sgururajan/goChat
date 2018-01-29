package db

import (
	"fmt"
	"goChat/Server/models"
	"goChat/Server/utils"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// MongoUserRepository - MongoUserRepository
type MongoUserRepository struct {
	session        *mgo.Session
	userCollection *mgo.Collection
}

const userCollectionName = "Users"

// NewMongoUserRepository - creates a new instance of MongoUserRepository
func NewMongoUserRepository(connString, dbName string) *MongoUserRepository {
	session, err := mgo.Dial(connString)
	if err != nil {
		utils.FailOnError(fmt.Errorf("Error while dialing Mongo: %s", err))
	}

	collection := session.DB(dbName).C(userCollectionName)
	client := MongoUserRepository{
		session:        session,
		userCollection: collection,
	}

	ensureIndexs(session, dbName)

	return &client
}

func ensureIndexs(s *mgo.Session, dbName string) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(dbName).C(userCollectionName)

	emailIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := c.EnsureIndex(emailIndex)
	failOnError(err)

	firstNameIndex := mgo.Index{
		Key:        []string{"firstName"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(firstNameIndex)
	failOnError(err)

	lastNameIndex := mgo.Index{
		Key:        []string{"lastName"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(lastNameIndex)
	failOnError(err)
}

// Interface members

// Create - Create
func (repo *MongoUserRepository) Create(user models.User) (string, error) {
	// check if the user already exist
	var tmpUser models.User
	err := repo.userCollection.Find(bson.M{"email": user.Email}).One(&tmpUser)

	if err == nil {
		return "", fmt.Errorf("User with email %s already exists", user.Email)
	}

	err = repo.userCollection.Insert(user)

	return "", nil
}

// GetUserByEmail - GetUserByEmail
func (repo *MongoUserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := repo.userCollection.Find(bson.M{"email": email}).One(&user)
	return user, err
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
