package main

import (
	"net/http"
	"strconv"
	"log"
	"github.com/nobelium/dalalstreet/controllers"
)

const (
	ADDRESS = ":3000"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/login", controllers.LoginProcessor).Methods("POST")
	r.HandleFunc("/logout", controllers.LogoutHandler).Methods("GET", "POST")
	r.HandleFunc("/", controllers.IndexHandler).Methods("GET")

	m := new(Middleware)
	m.UseCSRF()
	m.UseLogger()

	http.Handle("/", m.Init(r))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))

	address := HOST + ":" + strconv.Itoa(PORT)
	log.Println("Starting server on: ", address)
	http.ListenAndServe(address, nil)
}