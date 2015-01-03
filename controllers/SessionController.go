package controllers

import (
	"net/http"
	"log"
	// "fmt"
	"github.com/nobelium/dalalstreet/config"
	"github.com/nobelium/dalalstreet/models"
)

func LoginHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached LoginHandler")
	config.Render(res, config.T("login.html"), nil)
}

func AuthHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached AuthHandler")
	
	username, password := req.FormValue("username"), req.FormValue("password")

	user, err := models.Validate(username, password)
	
	if err != nil {
		log.Println("Failed to authenticate : ", username)
	}

	session := config.GetSession(req)
	var redirect_uri string
	if user == nil {
		session.AddFlash("User name or password is incorrect", config.MessageName)
		session.Save(req, res)
		redirect_uri = "/login"
	} else {
		session.Values["user"] = user
		redirect_uri, _ = session.Values["RedirectURI"].(string)
	}
	session.Save(req, res)
	http.Redirect(res, req, redirect_uri, http.StatusFound)
}

func LogoutHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached LogoutHandler")
	// fmt.Fprint(res, "Logged out")
	session := config.GetSession(req)
	session.AddFlash("You have been Logged out", config.MessageName)
	session.Save(req, res)

	http.Redirect(res, req, "/", http.StatusFound)
}