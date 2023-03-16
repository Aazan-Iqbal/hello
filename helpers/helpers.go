package helpers

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, page string) {
	// see a webpage
	// we will be using templates which are html mixed with dynamic data
	ts, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
