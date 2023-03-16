// Filename: cmd/web/middleware.go
package main

import (
	"net/http"
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
