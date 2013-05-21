package main

import (
	"net/http"
	"log"
)

func main() {
	log.Println("Starting Server")
	http.Handle("/", http.FileServer(http.Dir("./public/")))

	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
