package main

import (
	"net/http"
	"log"
	"reflect"
	"code.google.com/p/xsrftoken"
	"github.com/gorilla/context"
	"github.com/nobelium/dalalstreet/config"
	"github.com/nobelium/dalalstreet/models"
)

type MiddlewareFunc func (http.ResponseWriter, *http.Request)

type Middleware struct {
	middlewares []MiddlewareFunc
}

func (m *Middleware) Init(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := make([]reflect.Value, 2)
		in[0] = reflect.ValueOf(w)
		in[1] = reflect.ValueOf(r)
		for _,f := range m.middlewares {
			reflect.ValueOf(f).Call(in)
		}

		h.ServeHTTP(w, r)
	})
}

func (m *Middleware) Use(f MiddlewareFunc){
	m.middlewares = append(m.middlewares, f)
}

// Pre configured Middlewares
func (m *Middleware) UseCSRF() {
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		xsrftoken.Generate("samplekey", "asdf", r.Method)
		//TODO find a way to inject into request
	})
}

func (m *Middleware) UseLogger() {
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\"%s\" \"%s\" \"%s\"", r.RemoteAddr, r.Method, r.URL)
	})
}

// Set the user in gorilla context if its available in the session
func (m *Middleware) UseContextSetter() {
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		session := config.GetSession(r)
		user := session.Values["user"]
		if (user != nil) {
			user = user.(*models.User)
			context.Set(r, config.LoggedInUser, user)
		}
		log.Println("Found user ", user)
	})
}
