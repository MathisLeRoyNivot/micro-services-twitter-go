package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Rating struct {
	Id string `json:"Id"`
	Twitter_id string `json:"twitter_id"`
	Rated_twitter_id string `json:"rated_twitter_id"`
	Note string `json:"note"`
}
func allRating(w http.ResponseWriter, r *http.Request) {
	rate := Rating{Id:"1", Twitter_id:"1", Rated_twitter_id:"2", Note:"5"}
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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://geoffrey:NR0ONP0dl9n1amLB@cluster0-h7zuk.mongodb.net/test?retryWrites=true&w=majority"))
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