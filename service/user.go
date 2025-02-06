package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

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
