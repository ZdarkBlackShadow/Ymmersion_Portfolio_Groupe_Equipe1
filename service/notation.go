package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Rating(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var listOfNotes []int
		var value string
		for _, id := range []string{"1", "2", "3", "4", "5", "6"} {
			value = r.FormValue("rating_" + id)
			if value == "" {
				http.Error(w, "Veuillez renseigner toutes les notes", http.StatusBadRequest)
				return
			} else {
				valeur, err := strconv.Atoi(r.FormValue("rating_" + id))
				if err != nil {
					log.Fatal(err)
				}
				listOfNotes = append(listOfNotes, valeur)
			}
		}
		advice := r.FormValue("advice")
		if advice == "" {
			http.Error(w, "Veuillez renseigner un conseil", http.StatusBadRequest)
			return
		}
		// Ajout de la note et du conseil dans la base de données
		DataToadd := Advice{Note: listOfNotes, Advice: advice}
		var total float64
		var count int

		filename := "./data/user.json"
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		var Data UserStruct
		err = json.Unmarshal(bytes, &Data)
		if err != nil {
			log.Fatal(err)
		}
		Data.AllAdvice = append(Data.AllAdvice, DataToadd)
		for _, note := range Data.AllAdvice {
			for _, value := range note.Note {
				total += float64(value)
				count++
			}
		}
		Data.Average = total / float64(count)
		// Enregistrement des données
		file, err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		dataBytes, err := json.Marshal(Data)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.Write(dataBytes)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	} else {
		http.Error(w, "Méthode HTTP non autorisée", http.StatusMethodNotAllowed)
	}
}
