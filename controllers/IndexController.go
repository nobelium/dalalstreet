package controllers

import (
	"net/http"
	"log"
	"github.com/nobelium/dalalstreet/config"
)

func IndexHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached IndexHandler")
	session, _ := config.Store.Get(req, "flash-session")
	session.Values["store"] = "asdfasdf--"
	session.Save(req, res)

	// http.ServeFile(res, req, "./views/index.html")
	a := config.T("index.html").Execute(res, map[string]interface{}{})
	return a
}