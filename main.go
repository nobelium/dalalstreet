package main

import (
	"net/http"
	"strconv"
	"log"
)

const (
	HOST	= "0.0.0.0"
	PORT	= 3000
)

func main() {
	r := Router

	m := new(Middleware)
	m.UseCSRF()
	m.UseLogger()
	m.UseContextSetter()

	http.Handle("/", m.Init(r))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))

	address := HOST + ":" + strconv.Itoa(PORT)
	log.Println("Starting server on: ", address)
	http.ListenAndServe(address, nil)
}