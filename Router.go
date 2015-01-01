package main

import (
	"github.com/gorilla/mux"
	"github.com/nobelium/dalalstreet/controllers"
)

var Router = mux.NewRouter()

func init () {
	Router.HandleFunc("/login", controllers.LoginHandler).Methods("GET")
	Router.HandleFunc("/login", controllers.AuthHandler).Methods("POST")
	Router.HandleFunc("/logout", controllers.LogoutHandler).Methods("GET", "POST")

	Router.HandleFunc("/", controllers.IndexHandler).Methods("GET")
}