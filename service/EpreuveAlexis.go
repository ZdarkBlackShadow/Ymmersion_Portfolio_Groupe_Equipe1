package service

import (
	"log"
	"net/http"
)

func EpreuveAlexis(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "alexis", ALLEpreuve[2])
	if err != nil {
		log.Fatal(err)
	}
}
