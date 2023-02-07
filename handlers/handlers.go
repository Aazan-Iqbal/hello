package handlers

import (
	"net/http"

	"github.com/Aazan-Iqbal/hello/helpers"
)

// greeting function handler
func Greeting(w http.ResponseWriter, r *http.Request) {

}

// homepage function handler
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	helpers.RenderTemplates(w, "./static/html/home.page.tmpl")

}

// about page function handler
func About(w http.ResponseWriter, r *http.Request) {
	// day := time.Now().Weekday()

}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

}
