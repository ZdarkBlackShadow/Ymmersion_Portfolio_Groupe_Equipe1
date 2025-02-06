package service

import (
	"log"
	"net/http"
)

func YoanHandle(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "yoan", nil)
	if err != nil {
		log.Fatal(err)
	}
}
