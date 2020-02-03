package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"models"
	"time"
)

type RatingsDAO struct {
	Db_username string
	Db_password string
	Db_address  string
	Database    string
}

var db *mongo.Database


const (
	COLLECTION = "ratings"
)

// Establish a connection to database
func (r *RatingsDAO) Connect() {
	fmt.Println("En connexion")
	sessionOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/test?retryWrites=true&w=majority", r.Db_username, r.Db_password, r.Db_address))
	session, err := mongo.Connect(context.TODO(), sessionOptions)
	if err != nil {
		log.Fatal(err)
	}
	db = session.Database(fmt.Sprintf("%s", r.Database))
	fmt.Println("Vous êtes connecté")
}

// Get all ratings
func (m *RatingsDAO) FindAll() ([]models.Rating, error) {
	var ratings []models.Rating
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var cursor, err = db.Collection(COLLECTION).Find(ctx, bson.M{})
	if err != nil {
		return ratings , err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		var rating models.Rating
		cursor.Decode(&rating)
		ratings = append(ratings, rating)
	}
	return ratings, err
}

// Find rating by id
func (m *RatingsDAO) FindById(id primitive.ObjectID) (models.Rating, error) {
	var rating models.Rating
	filter := bson.M{"_id": id}
	err := db.Collection(COLLECTION).FindOne(context.TODO(), filter).Decode(&rating)
	return rating, err
}

// Insert rating
func (r *RatingsDAO) Insert(rating models.Rating) error {
	_, err := db.Collection(COLLECTION).InsertOne(context.TODO(), rating)
	return err
}

// Delete rating
func (r *RatingsDAO) Delete(rating models.Rating) error {
	_, err := db.Collection(COLLECTION).DeleteOne(context.TODO(), rating)
	return err
}
