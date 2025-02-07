package service

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RequestPasswd struct {
	Text string `json:"passwd"`
}

type ResponsePasswd struct {
	Text string `json:"text"`
}

func CheckStrengthPasswd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	var reqData RequestPasswd
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Erreur lors du décodage", http.StatusBadRequest)
		return
	}
	var resData ResponsePasswd
	if len(reqData.Text) < 8 {
		resData = ResponsePasswd{Text: "Le mot de passe est trop court ! (8 caractères minimum)"}
	} else {
		score := 0
		if strings.ContainsAny(reqData.Text, "0123456789") {
			score++
		}
		if strings.ContainsAny(reqData.Text, "abcdefghijklmnopqrstuvwxyz") {
			score++
		}
		if strings.ContainsAny(reqData.Text, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			score++
		}
		if strings.ContainsAny(reqData.Text, "!@#$%^&*()-_=+[]{}|;:,.<>/?") {
			score++
		}
		switch score {
		case 1:
			resData = ResponsePasswd{Text: "Le mot de passe est faible"}
		case 2:
			resData = ResponsePasswd{Text: "Le mot de passe est moyen"}
		default:
			resData = ResponsePasswd{Text: "Le mot de passe est fort"}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}
