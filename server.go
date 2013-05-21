package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func KittensHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"kittens": [
		{"id": 1, "name": "Bobby", "picture": "http://placekitten.com/200/200"},
		{"id": 2, "name": "Wally", "picture": "http://placekitten.com/200/200"},
	]}`))
}

func main() {
	log.Println("Starting Server")

	r := mux.NewRouter()
	r.HandleFunc("/api/kittens", KittensHandler).Methods("GET")
	http.Handle("/api/", r)

	http.Handle("/", http.FileServer(http.Dir("./public/")))

	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
