package dao

import (
	. "github.com/mlabouardy/movies-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type RatingsDAO struct {
	Server   string
	Database string
}
