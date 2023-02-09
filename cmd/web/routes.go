package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	// create a multiplexer
	// Used to store which handlers to call when a page requests an endpoint
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/greeting", app.Greeting)
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/message/create", app.MessageCreate)
	return mux
}
