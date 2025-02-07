package service

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Structure pour la requête d'analyse de texte
type AnalyseRequest struct {
	Text string `json:"inputText"`
}

// Structure pour la réponse d'analyse de texte
type AnalyseResponse struct {
	WordCount   int    `json:"wordCount"`
	CharCount   int    `json:"charCount"`
	LongestWord string `json:"longestWord"`
}

// Handler pour analyser un texte
func TextAnalyser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	var reqData AnalyseRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Erreur lors du décodage", http.StatusBadRequest)
		return
	}
	reqData.Text = strings.TrimSpace(reqData.Text)
	words := strings.Fields(reqData.Text)
	numWords := len(words)
	numChars := 0
	longestWord := ""
	for _, word := range words {
		numChars += len(word)
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	resData := AnalyseResponse{WordCount: numWords, CharCount: numChars, LongestWord: longestWord}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}
