package models

import "gopkg.in/mgo.v2/bson"

type Rating struct {
	Id               bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Twitter_id       string        `json:"twitter_id,omitempty" bson:"twitter_id,omitempty"`
	Rated_twitter_id string        `json:"rated_twitter_id,omitempty" bson:"rated_twitter_id,omitempty"`
	Note             string        `json:"note,omitempty" bson:"note,omitempty"`
}
