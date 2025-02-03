package service

import (
	"log"
	"net/http"
)

func EpreuveAdrien(w http.ResponseWriter, r *http.Request) {
	data, errjson := GetAllEpreuveData()
	if errjson != nil {
		log.Fatal(errjson)
	}
	err := Templates.ExecuteTemplate(w, "Adrien", data[1])
	if err != nil {
		log.Fatal(err)
	}
}
