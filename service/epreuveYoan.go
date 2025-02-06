package service

import (
	"log"
	"net/http"
)

func YoanHandle(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "yoan", ALLEpreuve[5])
	if err != nil {
		log.Fatal(err)
	}
}
