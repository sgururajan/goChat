package mongo

import (
	"fmt"
	"goChat/Server/models"
	"goChat/Server/utils"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// ConversationRepository - IConversationRepository implementation using MongoDB
type ConversationRepository struct {
	session    *mgo.Session
	collection *mgo.Collection
}

const conversationCollection = "conversations"

// NewConversationRepository - func to create new instance of MongoConversationRepository
func NewConversationRepository(connString, dbName string) *ConversationRepository {
	msession, err := mgo.Dial(connString)
	if err != nil {
		utils.FailOnError(fmt.Errorf("Error while dialing Mongo: %s", err))
	}

	mcollection := msession.DB(dbName).C(conversationCollection)
	client := &ConversationRepository{
		session:    msession,
		collection: mcollection,
	}

	client.ensureIndexes(msession, dbName)

	return client
}

func (repo *ConversationRepository) ensureIndexes(s *mgo.Session, dbName string) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(dbName).C(conversationCollection)
	partiesIndex := mgo.Index{
		Key:        []string{"participants"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := c.EnsureIndex(partiesIndex)
	utils.FailOnError(err)
}

// GetConversationsForUserID -  GetConversationsForUserID
func (repo *ConversationRepository) GetConversationsForUserID(userID string) ([]models.Conversation, error) {
	var conversation []models.Conversation
	err := repo.collection.Find(bson.M{"participants": userID}).All(&conversation)
	if err != nil {
		return conversation, fmt.Errorf("Error while fetching conversation for participants: %s", err)
	}

	return conversation, nil
}

// AddConversation - AddConversation
func (repo *ConversationRepository) AddConversation(conversation models.Conversation) (string, error) {
	cID := NewObjectID()
	conversation.ID = cID
	err := repo.collection.Insert(conversation)
	return "", err
}

// AddMessage - AddMessage
func (repo *ConversationRepository) AddMessage(message models.Message, convID string) error {
	pushMsg := bson.M{"$push": bson.M{"messages": message}}
	conv := bson.M{"_id": convID}

	err := repo.collection.Update(conv, pushMsg)
	return err
}

// GetConversationByID - GetConversationByID
func (repo *ConversationRepository) GetConversationByID(id string) (models.Conversation, error) {
	var conversation models.Conversation
	err := repo.collection.Find(bson.M{"_id": id}).One(&conversation)
	return conversation, err
}
