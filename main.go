package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"./controllers"
)

const (
	ADDRESS = ":3000"
)

func main() {
	log.Println("Starting server on: ", ADDRESS)

	r := mux.NewRouter()

	
	r.HandleFunc("/login", controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/login", controllers.LoginProcessor).Methods("POST")
	r.HandleFunc("/logout", controllers.LogoutHandler).Methods("GET", "POST")
	
	r.HandleFunc("/", controllers.IndexHandler).Methods("GET")

	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))

	http.ListenAndServe(ADDRESS, nil)
}