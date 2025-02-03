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

func Tableaudebord(w http.ResponseWriter, r *http.Request) {
	Datatdb := DataTDB{
		Datatdb: "Tableau de bord",
	}
	err := Templates.ExecuteTemplate(w, "tableaudebord", Datatdb)
	if err != nil {
		log.Fatal(err)
	}
}
