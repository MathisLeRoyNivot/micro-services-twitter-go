package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Rating struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Twitter_id       string        `json:"twitter_id,omitempty" bson:"twitter_id,omitempty"`
	Rated_twitter_id string        `json:"rated_twitter_id,omitempty" bson:"rated_twitter_id,omitempty"`
	Note             string        `json:"note,omitempty" bson:"note,omitempty"`
}
