package service

import (
	"log"
	"net/http"
)

func EpreuveJonathan(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "jonathan", ALLEpreuve[4])
	if err != nil {
		log.Fatal(err)
	}
}

func TextAnalyser(w http.ResponseWriter, r *http.Request) {

}