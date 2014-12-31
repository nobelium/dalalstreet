package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"./controllers"
	"log"
)

const (
	ADDRESS = ":3000"
)

func main() {
	log.Println("Starting server on: ", ADDRESS)

	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)

	http.Handle("/", r)

	http.ListenAndServe(ADDRESS, nil)
}