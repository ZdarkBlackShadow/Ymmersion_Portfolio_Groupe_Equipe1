package service

import (
	"log"
	"net/http"
)

func EpreuveKellyan(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "kellyan", ALLEpreuve[3])
	if err != nil {
		log.Fatal(err)
	}
}
