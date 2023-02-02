package handlers

import (
	"fmt"
	"net/http"
	"time"
	
)

// greeting function handler
func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to My Page!"))
}

// homepage function handler
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Wassaaaap"))
}

// about page function handler
func About(w http.ResponseWriter, r *http.Request) {
	day := time.Now().Weekday()
	w.Write([]byte(fmt.Sprintf("Have a good %s.", day)))
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("message created...."))
}
