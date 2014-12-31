package controllers

import (
	"net/http"
	"log"
)

func IndexHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached IndexHandler")
	http.ServeFile(res, req, "./views/index.html")
}