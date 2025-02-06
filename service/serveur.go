package service

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var Templates *template.Template
var ALLEpreuve []Exercice

func seq(start, end int) []int {
	var seq []int
	for i := start; i <= end; i++ {
		seq = append(seq, i)
	}
	return seq
}

func InitServer() {
	var err error
	//recuperer toutes les épreuves
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
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/exemple", Exemple)
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
	http.HandleFunc("/rating", Rating)
	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Initialisation du serveur
	fmt.Println("http://localhost:8080/home")
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
