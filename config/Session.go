package config

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var (
	Store = sessions.NewCookieStore([]byte("324546fa343e8b9067bb412d678a89e83629ffe23940"))
	SessionName = "flash-session"
	MessageName = "Message"
)

func GetSession (req *http.Request) *sessions.Session {
	session, err := Store.Get(req, SessionName)
	if err != nil {
		log.Fatalf("Failed to retreive session")
	}
	return session
}
