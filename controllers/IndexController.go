package controllers

import (
	"net/http"
	"log"
	"github.com/nobelium/dalalstreet/config"
	"html/template"
)

func IndexHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached IndexHandler")
	session, _ := config.Store.Get(req, "flash-session")
	session.Values["store"] = "asdfasdf--"
	session.Save(req, res)

	// http.ServeFile(res, req, "./views/index.html")
	
	t := template.Must(template.ParseFiles("views/_base.html", "views/index.html"))
	err := t.Execute(res, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}