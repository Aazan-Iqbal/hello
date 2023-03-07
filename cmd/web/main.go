// Filename: hello.go

package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Aazan-Iqbal/hello/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Create a new type
// Dependency Injection(DI):
// It is a way to neatly expose data to all the handlers. Acts as a centralized repository for
// all handlers to access data
type application struct {
	question models.QuestionModel
}

func main() {

	// creating a flag for specifying the port number when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	//grab our environment variable for our database from the .profile using the os package to interact with the OS.
	// We also add a help message for when an invalid dsn is passed
	dsn := flag.String("dsn", os.Getenv("RCSYSTEM_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	// create an instance of a connection pool(Pool of many reusable connections to the DB)
	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}

	// Create an instance of the application type
	app := &application{
		question: models.QuestionModel{DB: db},
	}
	defer db.Close()
	log.Println("Database connection pool established.")

	//Create a customized server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	// create our server
	log.Printf("Starting server on port %s", *addr) // print to show an attempt was made to start the server
	err = srv.ListenAndServe()                      // start the server in port 4000 and pass any errors to "err"

	log.Fatal(err) // should not be reached. Prints out errors if the server did not start properly

}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// let's ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
