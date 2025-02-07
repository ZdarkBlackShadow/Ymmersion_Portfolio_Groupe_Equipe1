package service

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Structure pour la requête de vérification de palindrome
type RequestPalindrome struct {
	Text string `json:"text"`
}

// Structure pour la réponse de vérification de palindrome
type ResponsePalindrome struct {
	Text string `json:"text"`
}

// Handler pour vérifier si un texte est un palindrome
func CheckPalindromeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	var reqData RequestPalindrome
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Erreur lors du décodage", http.StatusBadRequest)
		return
	}
	s := strings.ToLower(strings.ReplaceAll(reqData.Text, " ", ""))
	n := len(s)
	isPalindrome := true
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			isPalindrome = false
			break
		}
	}
	var resData ResponsePalindrome
	if isPalindrome {
		resData = ResponsePalindrome{Text: "✅ C'est un palindrome ! 🎉"}
	} else {
		resData = ResponsePalindrome{Text: "❌ Ce n'est pas un palindrome."}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}
