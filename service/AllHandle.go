package service

import (
	"log"
	"net/http"
)

// Handler pour la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Data := GetUser()
	ereur := Templates.ExecuteTemplate(w, "home", Data)
	if ereur != nil {
		log.Fatal(ereur)
	}
}

// Handler pour l'épreuve d'Alexandre
func EpreuveAlexandre(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "Alexandre", ALLEpreuve[0])
	if err != nil {
		log.Fatal(err)
	}
}

// Handler pour l'épreuve d'Adrien
func EpreuveAdrien(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "Adrien", ALLEpreuve[1])
	if err != nil {
		log.Fatal(err)
	}
}

// Handler pour l'épreuve d'Alexis
func EpreuveAlexis(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "alexis", ALLEpreuve[2])
	if err != nil {
		log.Fatal(err)
	}
}

// Handler pour l'épreuve de Jonathan
func EpreuveJonathan(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "jonathan", ALLEpreuve[4])
	if err != nil {
		log.Fatal(err)
	}
}

// Handler pour l'épreuve de Kellyan
func EpreuveKellyan(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "kellyan", ALLEpreuve[3])
	if err != nil {
		log.Fatal(err)
	}
}

// Handler pour l'épreuve de Yoan
func YoanHandle(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "yoan", ALLEpreuve[5])
	if err != nil {
		log.Fatal(err)
	}
}

// Handler pour le tableau de bord
func Tableaudebord(w http.ResponseWriter, r *http.Request) {
	Data := GetUser()
	err := Templates.ExecuteTemplate(w, "tableaudebord", Data)
	if err != nil {
		log.Fatal(err)
	}
}
