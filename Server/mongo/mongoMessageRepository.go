package mongo

import (
	"goChat/Server/models"
	"goChat/Server/utils"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// MessageRepository - Mongo implementation for IMessageRepository
type MessageRepository struct {
	session    *mgo.Session
	collection *mgo.Collection
}

const messageCollectionName = "Messages"

// NewMessageRepository - Creates new instance of MessageRepository
func NewMessageRepository(connString, dbName string) *MessageRepository {
	mSession, err := mgo.Dial(connString)
	if err != nil {
		utils.FailOnError(err)
	}

	mCollection := mSession.DB(dbName).C(messageCollectionName)
	client := &MessageRepository{
		session:    mSession,
		collection: mCollection,
	}

	client.ensureIndexes()

	return client
}

func (repo *MessageRepository) ensureIndexes() {
	session := repo.session.Copy()
	defer session.Close()

	conversationIndex := mgo.Index{
		Key:        []string{"conversationID"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := repo.collection.EnsureIndex(conversationIndex)
	utils.FailOnError(err)
}

// GetMessagesByConversation - GetMessagesByConversation
func (repo *MessageRepository) GetMessagesByConversation(conversationID string, page, count int) ([]models.Message, error) {
	var messages []models.Message
	err := repo.collection.Find(bson.M{"conversationID": conversationID}).Sort("-timeStamp").Skip((page - 1) * count).Limit(count).All(&messages)
	return messages, err
}

// AddMessage - Add message
func (repo *MessageRepository) AddMessage(message models.Message) error {
	message.MessageID = NewObjectId()
	err := repo.collection.Insert(message)
	return err
}
