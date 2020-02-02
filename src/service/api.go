package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"

	. "github.com/MathisLeRoyNivot/micro-services-twitter-go/src/service/dao"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type Rating struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Twitter_id       string             `json:"twitter_id,omitempty" bson:"twitter_id,omitempty"`
	Rated_twitter_id string             `json:"rated_twitter_id,omitempty" bson:"rated_twitter_id,omitempty"`
	Note             string             `json:"note,omitempty" bson:"note,omitempty"`
}

/*func allRating(w http.ResponseWriter, r *http.Request) {
	rate := Rating{Id: "1", Twitter_id: "1", Rated_twitter_id: "2", Note: "5"}
	fmt.Println("Voici une note pour ce tweet")
	json.NewEncoder(w).Encode(rate)
}

func testPostRating(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/ratings", allRating).Methods("GET")
	myRouter.HandleFunc("/ratings", testPostRating).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}

func main() {
	// godotenv package
	username := goDotEnvVariable("DB_USERNAME")
	password := goDotEnvVariable("DB_PASSWORD")
	dbAddress := goDotEnvVariable("DB_ADDRESS")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/test?retryWrites=true&w=majority", username, password, dbAddress)))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Vous êtes connecté à la base")
	collection := client.Database("twitter_microservices_api").Collection("ratings")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(collection)
	handleRequests()
	client.Disconnect(context.TODO())
} */

var client *mongo.Client

func allRating(w http.ResponseWriter, r *http.Request) {
	rate := Rating{Twitter_id: "1", Rated_twitter_id: "2", Note: "5"}
	fmt.Println("Voici une note pour ce tweet")
	json.NewEncoder(w).Encode(rate)
}

func addOneRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var rating Rating
	json.NewDecoder(r.Body).Decode(&rating)
	collection := client.Database("twitter_microservices_api").Collection("ratings")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	result, _ := collection.InsertOne(ctx, rating)
	json.NewEncoder(w).Encode(result)
}

func getRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var ratings []Rating
	collection := client.Database("twitter_microservices_api").Collection("ratings")
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message"}` + err.Error() + `"}"`))
		return
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message"}` + err.Error() + `"}"`))
		return
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var rating Rating
		cursor.Decode(&rating)
		ratings = append(ratings, rating)
	}
	json.NewEncoder(w).Encode(ratings)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/rating", allRating).Methods("GET")
	myRouter.HandleFunc("/ratings", getRating).Methods("GET")
	myRouter.HandleFunc("/rating", addOneRate).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}

func main() {
	// godotenv package
	username := goDotEnvVariable("DB_USERNAME")
	password := goDotEnvVariable("DB_PASSWORD")
	dbAddress := goDotEnvVariable("DB_ADDRESS")

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/test?retryWrites=true&w=majority", username, password, dbAddress)))
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://geoffrey:NR0ONP0dl9n1amLB@cluster0-h7zuk.mongodb.net/test?retryWrites=true&w=majority"))
	_ = client.Ping(context.TODO(), nil)
	fmt.Println("Vous êtes connecté à la base")
	handleRequests()
	client.Disconnect(context.TODO())
}
