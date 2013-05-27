package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"strconv"
)

type Kitten struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type KittenJSON struct {
	Kitten Kitten `json:"kitten"`
}

type KittensJSON struct {
	Kittens []Kitten `json:"kittens"`
}

var idCounter int
var kittens []Kitten

func KittensHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(KittensJSON{Kittens: kittens})
	if err != nil {
		panic(err)
	}
	w.Write(j)
}

func CreateKittenHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming kitten from the request body
	var kittenJSON KittenJSON
	err := json.NewDecoder(r.Body).Decode(&kittenJSON)
	if err != nil {
		panic(err)
	}

	idCounter++

	// Grab the kitten and set some dummy data
	kitten := kittenJSON.Kitten
	kitten.Id = idCounter
	kitten.Picture = "http://placekitten.com/300/200"

	kittens = append(kittens, kitten)

	// Serialize the modified kitten to JSON
	j, err := json.Marshal(KittenJSON{Kitten: kitten})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func UpdateKittenHandler(w http.ResponseWriter, r *http.Request) {
	// Grab the kitten's id from the incoming url
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	// Decode the incoming kitten json
	var kittenJSON KittenJSON
	err = json.NewDecoder(r.Body).Decode(&kittenJSON)
	if err != nil {
		panic(err)
	}

	// Find the kitten in our kittens slice and upate it's name
	for index, _ := range kittens {
		if kittens[index].Id == id {
			kittens[index].Name = kittenJSON.Kitten.Name
		}
	}

	// Respond with a 204 indicating success, but no content
	w.WriteHeader(http.StatusNoContent)
}

func DeleteKittenHandler(w http.ResponseWriter, r *http.Request) {
	// Grab the kitten's id from the incoming url
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	// Find the index of the kitten
	kittenIndex := -1
	for index, _ := range kittens {
		if kittens[index].Id == id {
			kittenIndex = index
			break
		}
	}

	// If we actually found a kitten remove it from the slice
	if kittenIndex != -1 {
		kittens = append(kittens[:kittenIndex], kittens[kittenIndex+1:]...)
	}

	// Respond with a 204 indicating success, but no content
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	log.Println("Starting Server")

	r := mux.NewRouter()
	r.HandleFunc("/api/kittens", KittensHandler).Methods("GET")
	r.HandleFunc("/api/kittens", CreateKittenHandler).Methods("POST")
	r.HandleFunc("/api/kittens/{id}", UpdateKittenHandler).Methods("PUT")
	r.HandleFunc("/api/kittens/{id}", DeleteKittenHandler).Methods("DELETE")
	http.Handle("/api/", r)

	http.Handle("/", http.FileServer(http.Dir("./public/")))

	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
