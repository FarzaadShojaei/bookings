package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit The Page")
		next.ServeHTTP(w, r)
	})
}

//adds CSRF protection to all POST requests
func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//loads and saves the session on every Request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
