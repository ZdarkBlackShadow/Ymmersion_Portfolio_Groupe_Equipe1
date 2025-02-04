package service

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var Templates *template.Template
var ALLEpreuve []Exercice

func InitServer() {
	var err error
	ALLEpreuve, err = GetAllEpreuveData()
	if err != nil {
		log.Fatal(err)
	}
	Templates, err = template.New("").ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	http.HandleFunc("/exemple", Exemple)
	http.HandleFunc("/tableaudebord", Tableaudebord)
	http.HandleFunc("/adrien", EpreuveAdrien)
	http.HandleFunc("/tableau_epreuve", PageEpreuve)
	http.HandleFunc("/alexandre", EpreuveAlexandre)
	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Initialisation du serveur
	fmt.Println("http://localhost:8080/exemple")
	http.ListenAndServe("localhost:8080", nil)
}
