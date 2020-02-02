package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
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
	Id               string `json:"Id"`
	Twitter_id       string `json:"twitter_id"`
	Rated_twitter_id string `json:"rated_twitter_id"`
	Note             string `json:"note"`
}

func allRating(w http.ResponseWriter, r *http.Request) {
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
	myRouter.HandleFunc("/rating", allRating).Methods("GET")
	myRouter.HandleFunc("/rating", testPostRating).Methods("POST")
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
}
