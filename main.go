package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)
}