package controllers

import (
	"net/http"
	"log"
	"github.com/nobelium/dalalstreet/config"
)

func IndexHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached IndexHandler")
	// session := config.GetSession(req)
	// session.Values["store"] = "asdfasdf--"
	// session.Save(req, res)

	// http.ServeFile(res, req, "./views/index.html")
	config.Render(res, config.T("index.html"), nil)
}