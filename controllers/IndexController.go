package controllers

import (
	"net/http"
	"log"
	"github.com/nobelium/dalalstreet/config"
)

func IndexHandler (res http.ResponseWriter, req *http.Request) {
	// http.ServeFile(res, req, "./views/index.html")
	log.Println("Reached IndexHandler")
	session := config.GetSession(req)
	user, _ := session.Values["user"]
	msg := session.Flashes()
	session.Save(req, res)

	config.Render(res, config.T("index.html"), map[string]interface{}{
			"user" : user,
			"flash" : msg,
		})
}