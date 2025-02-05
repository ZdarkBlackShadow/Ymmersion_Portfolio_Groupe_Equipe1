package service

import (
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ereur := Templates.ExecuteTemplate(w, "home", nil)
	if ereur != nil {
		log.Fatal(ereur)
	}
}
