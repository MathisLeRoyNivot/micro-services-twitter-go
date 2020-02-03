package main

import (
	"config"
	"dao"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"models"
	"net/http"
)

var account = config.Config{}
var data = dao.RatingsDAO{}

func AllRatingsEndPoint(w http.ResponseWriter, r *http.Request) {
	ratings, err := data.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, ratings)
}
func FindRatingEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id , _ := primitive.ObjectIDFromHex(params["id"])
	rating, err := data.FindById(id)
	fmt.Println(rating)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Rating ID")
		return
	}
	respondWithJson(w, http.StatusOK, rating)
}
func CreateRatingEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var rating models.Rating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	rating.Id = primitive.NewObjectID()
	if err := data.Insert(rating); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, rating)
}
func DeleteRatingEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var rating models.Rating
	if err := json.NewDecoder(r.Body).Decode(&rating);
	err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := data.Delete(rating);
	err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string] string {
		"result": "success",
	})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	account.Read()
	fmt.Println("cherche une femme")
	data.Db_username = account.Db_username
	data.Db_password = account.Db_password
	data.Db_address = account.Db_address
	data.Database = account.Database
	data.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ratings", AllRatingsEndPoint).Methods("GET")
	r.HandleFunc("/rating", CreateRatingEndPoint).Methods("POST")
	r.HandleFunc("/rating", DeleteRatingEndPoint).Methods("DELETE")
	r.HandleFunc("/ratings/{id}", FindRatingEndpoint).Methods("GET")
	err := http.ListenAndServe(":8002", r)
	if err != nil {
		log.Fatal(err)
	}
}
