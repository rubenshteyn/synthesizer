package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("secret-key"))
)

func CreateSession(w http.ResponseWriter, r *http.Request, name string) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}

	session.Values["name"] = name
	session.Values["authenticated"] = true
	session.Options.MaxAge = 300
	return session.Save(r, w)
}

func IsAuthenticated(r *http.Request) bool {
	session, _ := store.Get(r, "session-name")
	auth, ok := session.Values["authenticated"].(bool)
	return ok && auth
}
