package service

import (
	"log"
	"net/http"
)

func EpreuveKellyan(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "kellyan", nil)
	if err != nil {
		log.Fatal(err)
	}
}
