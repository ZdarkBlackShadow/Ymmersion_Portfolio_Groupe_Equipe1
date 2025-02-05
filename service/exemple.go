package service

import (
	"log"
	"net/http"
)

func Exemple(w http.ResponseWriter, r *http.Request) {
	type Da struct {
		Data Exercice
	}
	Data, err := GetAllEpreuveData()
	if err != nil {
		log.Fatal(err)
	}
	Datastruct := Da{
		Data: Data[0],
	}
	err = Templates.ExecuteTemplate(w, "exemple", Datastruct)
	if err != nil {
		log.Fatal(err)
	}
}

func Tableaudebord(w http.ResponseWriter, r *http.Request) {
	Data := GetUser()
	err := Templates.ExecuteTemplate(w, "tableaudebord", Data)
	if err != nil {
		log.Fatal(err)
	}
}
