// Filename: hello.go

package main

import (
	"flag"
	"log"
	"net/http"
)

// Create a new type
type application struct{}

func main() {

	// creating a flag for specifying the port number when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	flag.Parse()

	// Create an instance of the application type
	app := &application{}

	//get the routes

	//Create a customized server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	// create our server
	log.Printf("Starting server on port %s", *addr) // print to show an attempt was made to start the server
	err := srv.ListenAndServe()                     // start the server in port 4000 and pass any errors to "err"

	log.Fatal(err) // should not be reached. Prints out errors if the server did not start properly

}
