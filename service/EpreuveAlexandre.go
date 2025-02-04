package service

import (
	"log"
	"net/http"
)

func EpreuveAlexandre(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "Alexandre", ALLEpreuve[0])
	if err != nil {
		log.Fatal(err)
	}
}
