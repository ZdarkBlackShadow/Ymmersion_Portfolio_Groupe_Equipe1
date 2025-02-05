package service

import (
	"log"
	"net/http"
)

func EpreuveAlexis(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "Alexis", ALLEpreuve[2])
	if err != nil {
		log.Fatal(err)
	}
}