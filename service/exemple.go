package service

import (
	"log"
	"net/http"
)

func Exemple(w http.ResponseWriter, r *http.Request) {
	Data := DataExample{
		Data: "Hello world!",
	}
	err := Templates.ExecuteTemplate(w, "exemple", Data)
	if err != nil {
		log.Fatal(err)
	}
}
