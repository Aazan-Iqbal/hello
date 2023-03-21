// Filename: cmd/web/middleware.go
package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

// Create a middleware function to intercept responses from the server to a client to set some secutiy headers
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			w.Header().Set("X-Frame-Options", "deny")

			next.ServeHTTP(w, r)
		})
}

func (app *application) logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// when the request comes to me
			start := time.Now()
			app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
			next.ServeHTTP(w, r)
			// when the response comes to em
			app.infoLog.Printf("Request took %v", time.Since(start))

		})
}

func (app *application) recoverPanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("connection", "closed")
					trace := fmt.Sprintf("%s \n", debug.Stack())
					app.errorLog.Output(2, trace)
					http.Error(w, http.StatusText(http.StatusInternalServerError),
						http.StatusInternalServerError)
				}
			}() // brackets for the defer()

			next.ServeHTTP(w, r)
		})
}
