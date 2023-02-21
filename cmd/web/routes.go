package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//router function that initializes our router to create a handler
func (app *application) routes() http.Handler {

	// create a multiplexer
	// Used to store which handlers to call when a page requests an endpoint
	router := httprouter.New()

	//create a file server - to store and serve static content
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/create", app.Greeting)
	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodGet, "/about", app.About)
	router.HandlerFunc(http.MethodPost, "/create", app.MessageCreate)

	return router
}
