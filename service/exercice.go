package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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
