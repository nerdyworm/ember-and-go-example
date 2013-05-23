package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
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

	// Grab the kitten and set some dummy data
	kitten := kittenJSON.Kitten
	kitten.Id = len(kittens) + 1
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

func main() {
	log.Println("Starting Server")

	r := mux.NewRouter()
	r.HandleFunc("/api/kittens", KittensHandler).Methods("GET")
	r.HandleFunc("/api/kittens", CreateKittenHandler).Methods("POST")
	http.Handle("/api/", r)

	http.Handle("/", http.FileServer(http.Dir("./public/")))

	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
