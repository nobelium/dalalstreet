package main

import (
	"net/http"
	"log"
)

const (
	ADDRESS = ":3000"
)

func main() {
	log.Println("Starting server on: ", ADDRESS)

	http.Handle("/", Router)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))

	http.ListenAndServe(ADDRESS, nil)
}