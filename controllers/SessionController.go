package controllers

import (
	"net/http"
	"log"
	"fmt"
)

func LoginHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached LoginHandler")
}

func LoginProcessor (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached LoginHandler")
}

func LogoutHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached LogoutHandler")
	fmt.Fprint(res, "Logged out")
}