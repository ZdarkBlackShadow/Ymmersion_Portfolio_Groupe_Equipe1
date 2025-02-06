package service

import (
	"log"
	"net/http"
)

func EpreuveJonathan(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "jonathan", nil)
	if err != nil {
		log.Fatal(err)
	}
}
