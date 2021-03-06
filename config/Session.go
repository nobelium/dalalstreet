package config

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var (
	Store = sessions.NewCookieStore([]byte("324546fa343e8b9067bb412d678a89e83629ffe23940"))
	CSRF_key = "asd89389hfwpiuhefhibjkas7338iu39898wf8d899fds9"
	SessionName = "flash-session"
	MessageName = "Message"
)

func init () {
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
	}
}

func GetSession (req *http.Request) *sessions.Session {
	session, err := Store.Get(req, SessionName)
	if err != nil {
		log.Fatalf("Failed to retreive session")
	}
	return session
}
