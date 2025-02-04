package service

import (
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	data, errjson := GetEpreuvePlusTest()
	if errjson != nil {
		log.Fatal(errjson)
	}
	err := Templates.ExecuteTemplate(w, "Adrien", data)
	if err != nil {
		log.Fatal(err)
	}
}

func GetEpreuvePlusTest() (TestExercice, error) {
data, errjson := GetAllEpreuveData()
    if errjson != nil {
        return TestExercice{}, errjson
    }
	var dataPage TestExercice
	dataPage.Exercice = data[1]
	if dataPage.OnGoing {
		
	}
}
