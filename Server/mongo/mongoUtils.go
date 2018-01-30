package mongo

import (
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func NewObjectId() string {
	return strings.ToUpper(bson.NewObjectId().Hex())
}
