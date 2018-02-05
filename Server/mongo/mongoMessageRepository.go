package mongo

import (
	"goChat/Server/models"
	"goChat/Server/utils"
	"time"

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
func (repo *MessageRepository) GetMessagesByConversation(conversationID string, skipCount, count int) ([]models.Message, error) {
	var messages []models.Message
	err := repo.collection.Find(bson.M{"conversationID": conversationID}).Sort("-timeStamp").Skip(skipCount).Limit(count).All(&messages)
	return messages, err
}

// AddMessage - Add message
func (repo *MessageRepository) AddMessage(message models.Message) error {
	message.MessageID = NewObjectID()
	message.TimeStamp = time.Now()
	err := repo.collection.Insert(message)
	return err
}

// UpdateMessageAsRead - UpdateMessageAsRead
func (repo *MessageRepository) UpdateMessageAsRead(msgID, participantID string) {
	findQuery := bson.M{"_id": msgID, "participantsState.participantID": participantID}
	updateQuery := bson.M{"$set": bson.M{"participants.$.IsRead": true}}
	repo.collection.Upsert(findQuery, updateQuery)
}
