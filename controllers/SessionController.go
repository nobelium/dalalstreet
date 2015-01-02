package controllers

import (
	"net/http"
	"log"
	// "fmt"
	"github.com/nobelium/dalalstreet/config"
)

func LoginHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached LoginHandler")
	config.Render(res, config.T("login.html"), nil)
}

func AuthHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached AuthHandler")
	http.Redirect(res, req, "/", http.StatusFound)
}

func LogoutHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached LogoutHandler")
	// fmt.Fprint(res, "Logged out")
	session := config.GetSession(req)
	session.AddFlash("You have been Logged out", config.MessageName)
	session.Save(req, res)

	http.Redirect(res, req, "/", http.StatusFound)
}