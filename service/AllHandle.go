package service

import (
	"log"
	"net/http"
)

/*
Ficher contenant tout les handler
*/

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Data := GetUser()
	ereur := Templates.ExecuteTemplate(w, "home", Data)
	if ereur != nil {
		log.Fatal(ereur)
	}
}

func EpreuveAlexandre(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "Alexandre", ALLEpreuve[0])
	if err != nil {
		log.Fatal(err)
	}
}

func EpreuveAdrien(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "Adrien", ALLEpreuve[1])
	if err != nil {
		log.Fatal(err)
	}
}

func EpreuveAlexis(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "alexis", ALLEpreuve[2])
	if err != nil {
		log.Fatal(err)
	}
}

func EpreuveJonathan(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "jonathan", ALLEpreuve[4])
	if err != nil {
		log.Fatal(err)
	}
}

func EpreuveKellyan(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "kellyan", ALLEpreuve[3])
	if err != nil {
		log.Fatal(err)
	}
}

func YoanHandle(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "yoan", ALLEpreuve[5])
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
