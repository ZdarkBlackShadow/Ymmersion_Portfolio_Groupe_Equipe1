package service

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// Structure pour la requête de génération de mot de passe
type RequestData struct {
	Length int `json:"length"`
}

// Structure pour la réponse de génération de mot de passe
type ResponseData struct {
	Mdp string `json:"mdp"`
}

// Handler pour générer un mot de passe
func GeneratePasswdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	var reqData RequestData
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Erreur lors du décodage", http.StatusBadRequest)
		return
	}
	if reqData.Length < 8 || reqData.Length > 16 {
		http.Error(w, "Longueur du mot de passe invalide", http.StatusBadRequest)
		return
	}
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digits := []rune("0123456789")
	var password []rune
	for len(password) < reqData.Length {
		password = append(password, letters[rand.Intn(len(letters))])
		password = append(password, digits[rand.Intn(len(digits))])
	}
	if len(password) > reqData.Length {
		password = password[:reqData.Length]
	}
	mdp := string(password)

	resData := ResponseData{Mdp: mdp}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}
