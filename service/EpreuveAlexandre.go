package service

import (
	"log"
	"net/http"
)

func EpreuveAlexandre(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "Alex", ALLEpreuve[0])
	if err != nil {
		log.Fatal(err)
	}
}
