package dao

import (
	. "github.com/MathisLeRoyNivot/micro-services-twitter-go/src/service/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type RatingsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "ratings"
)

// Establish a connection to database
func (r *RatingsDAO) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Get all ratings
func (m *RatingsDAO) FindAll() ([]Rating, error) {
	var ratings []Rating
	err := db.C(COLLECTION).Find(bson.M{}).All(&ratings)
	return ratings, err
}

// Find rating by id
func (m *RatingsDAO) FindById(id string) (Rating, error) {
	var rating Rating
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&rating)
	return rating, err
}

// Insert rating
func (r *RatingsDAO) Insert(movie Rating) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete rating
func (r *RatingsDAO) Delete(rating Rating) error {
	err := db.C(COLLECTION).Remove(&rating)
	return err
}
