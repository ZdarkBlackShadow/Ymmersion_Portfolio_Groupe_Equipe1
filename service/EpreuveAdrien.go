package service

import (
	"log"
	"net/http"
)

func EpreuveAdrien(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "Adrien", ALLEpreuve[1])
	if err != nil {
		log.Fatal(err)
	}
}
