package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Fonction pour récupérer toutes les données des épreuves
func GetAllEpreuveData() ([]Exercice, error) {
	filename := "./data/epreuve.json"
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var exercices []Exercice
	err = json.Unmarshal(bytes, &exercices)
	if err != nil {
		return nil, err
	}

	return exercices, nil
}

// Fonction pour charger les événements
func loadEvents() (WeekSchedule, error) {
	file, err := os.Open("./data/events.json")
	if err != nil {
		return WeekSchedule{}, err
	}
	defer file.Close()

	var schedule WeekSchedule
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&schedule)
	if err != nil {
		return WeekSchedule{}, err
	}

	return schedule, nil
}

// Fonction pour récupérer les données des utilisateurs
func GetUser() UserStruct {
	filePath := "./data/user.json"
	var data UserStruct
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier : %v", err)
	}
	if len(fileContent) > 0 {
		if err := json.Unmarshal(fileContent, &data); err != nil {
			log.Fatalf("Erreur lors du décodage JSON : %v", err)
		}
	}
	return data
}
