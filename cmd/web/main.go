// Filename: hello.go

package main

import (
	"log"
	"net/http"

	"github.com/aazaniqbal/hello/handlers"
)

func main() {

	// create a multiplexer
	// Used to store which handlers to call when a page requests an endpoint
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/greeting", handlers.Greeting)
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/message/create", handlers.MessageCreate)

	// create our server
	log.Println("Starting server on port :4000") // print to show an attempt was made to start the server
	err := http.ListenAndServe(":4000", mux)     // start the server in port 4000 and pass any errors to "err"

	log.Fatal(err) // should not be reached. Prints out errors if the server did not start properly

}
