// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))
	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave)

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/about", app.about)
	router.HandlerFunc(http.MethodGet, "/poll/reply", app.pollReplyShow)
	router.Handler(http.MethodPost, "/poll/reply", dynamicMiddleware.ThenFunc(app.pollReplySubmit))
	router.Handler(http.MethodGet, "/poll/success", dynamicMiddleware.ThenFunc(app.pollSuccessShow))
	router.HandlerFunc(http.MethodGet, "/poll/create", app.pollCreateShow)
	router.Handler(http.MethodPost, "/poll/create", dynamicMiddleware.ThenFunc(app.pollCreateSubmit))
	router.HandlerFunc(http.MethodGet, "/options/create", app.optionsCreateShow)
	router.HandlerFunc(http.MethodPost, "/options/create", app.optionsCreateSubmit)

	// tidy up middleware chain
	// ones at the bottom/left are more deeply nested than the ones on the top/right
	standardMiddleware := alice.New(app.recoverPanicMiddleware,
		app.logRequestMiddleware,
		SecurityHeadersMiddleware,
	)

	//returns the router to our middleware before it hits the server and goes to a client to check the
	// contents and append things to the contents or block them if necessary.
	return standardMiddleware.Then(router)
}
