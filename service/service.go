package service

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var Templates *template.Template

func InitServer() {
	var err error
	Templates, err = template.New("").ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	http.HandleFunc("/exemple", Exemple)
	http.HandleFunc("/tableaudebord", Tableaudebord)
	http.HandleFunc("/adrien", EpreuveAdrien)
	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Initialisation du serveur
	fmt.Println("http://localhost:8080/exemple")
	http.ListenAndServe("localhost:8080", nil)
}
