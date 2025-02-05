package service

import (
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Data := GetUser()
	ereur := Templates.ExecuteTemplate(w, "home", Data)
	if ereur != nil {
		log.Fatal(ereur)
	}
}
