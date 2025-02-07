package service

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Variables globales pour les templates et les épreuves
var Templates *template.Template
var ALLEpreuve []Exercice

// Fonction pour convertir une note en séquence
func seq(nb float64) int {
	if nb >= 0 && nb < 0.5 {
		return 0
	} else if nb >= 0.5 && nb < 1 {
		return 1
	} else if nb >= 1 && nb < 1.5 {
		return 2
	} else if nb >= 1.5 && nb < 2 {
		return 3
	} else if nb >= 2 && nb < 2.5 {
		return 4
	} else if nb >= 2.5 && nb < 3 {
		return 5
	} else if nb >= 3 && nb < 3.5 {
		return 6
	} else if nb >= 3.5 && nb < 4 {
		return 7
	} else if nb >= 4 && nb < 4.5 {
		return 8
	} else if nb >= 4.5 && nb < 5 {
		return 9
	} else {
		return 10
	}
}

// Fonction d'initialisation du serveur
func InitServer() {
	var err error
	// Récupérer toutes les épreuves
	ALLEpreuve, err = GetAllEpreuveData()
	if err != nil {
		log.Fatal(err)
	}
	funcMap := template.FuncMap{
		"seq": seq,
	}
	Templates, err = template.New("").Funcs(funcMap).ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	// Définir les handlers pour les différentes routes
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/tableaudebord", Tableaudebord)
	http.HandleFunc("/adrien", EpreuveAdrien)
	http.HandleFunc("/alexandre", EpreuveAlexandre)
	http.HandleFunc("/calendar", calendarHandler)
	http.HandleFunc("/alexis", EpreuveAlexis)
	http.HandleFunc("/generatePasswd", GeneratePasswdHandler)
	http.HandleFunc("/kellyan", EpreuveKellyan)
	http.HandleFunc("/checkPalindrome", CheckPalindromeHandler)
	http.HandleFunc("/jonathan", EpreuveJonathan)
	http.HandleFunc("/analyzeText", TextAnalyser)
	http.HandleFunc("/yoan", YoanHandle)
	http.HandleFunc("/checkPasswd", CheckStrengthPasswd)
	http.HandleFunc("/rating", Rating)
	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	// Initialisation du serveur
	fmt.Println("http://localhost:8080/home")
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
