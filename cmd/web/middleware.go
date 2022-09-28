package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// func WritetoConsole(next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Hit The Page")
// 		next.ServeHTTP(w, r)
// 	})
// }

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// * SetBaseCookie to make sure that the token it generates is available on a per page basis
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",   // * To Be Applied to the entire site
		Secure:   false, // * because NOW we're not running on https
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
