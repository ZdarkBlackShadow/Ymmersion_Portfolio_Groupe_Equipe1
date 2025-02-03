package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetAllEpreuveData() ([]Exercice, error) {
	filename := "./data/epreuve.json"
	// Ouvrir le fichier
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Lire le contenu du fichier
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parser le JSON dans une slice d'Exercice
	var exercices []Exercice
	err = json.Unmarshal(bytes, &exercices)
	if err != nil {
		return nil, err
	}

	return exercices, nil
}
