package mongo

import (
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// NewObjectID - creates new bson objectid
func NewObjectID() string {
	return strings.ToUpper(bson.NewObjectId().Hex())
}
