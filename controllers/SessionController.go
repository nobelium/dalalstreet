package controllers

import (
	"net/http"
	"log"
	"encoding/gob"
	"github.com/nobelium/dalalstreet/config"
	"github.com/nobelium/dalalstreet/models"
)

func init () {
	gob.Register(&models.User{})
}

func LoginHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached LoginHandler")

	if(models.IsLoggedIn(req)) {
		http.Redirect(res, req, "/", http.StatusFound)
	}

	session := config.GetSession(req)
	session.Values["RedirectURI"] = req.URL.Query().Get("redirect_uri")
	session.Save(req, res)

	config.Render(res, config.T("login.html"), map[string]interface{}{
			"moreStyles" : [...]string{"login.css"},
		})
}

func AuthHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached AuthHandler")
	
	if(models.IsLoggedIn(req)) {
		http.Redirect(res, req, "/", http.StatusFound)
	}
	
	username, password := req.FormValue("Username"), req.FormValue("Password")

	user, err := models.Validate(username, password)
	session := config.GetSession(req)
	var redirect_uri string
	if err != nil || user == nil {
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
	delete(session.Values, "user")
	delete(session.Values, "RedirectURI")

	session.Save(req, res)

	http.Redirect(res, req, "/", http.StatusFound)
}