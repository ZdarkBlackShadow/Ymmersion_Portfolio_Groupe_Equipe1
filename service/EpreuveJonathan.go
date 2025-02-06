package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func EpreuveJonathan(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "jonathan", ALLEpreuve[4])
	if err != nil {
		log.Fatal(err)
	}
}

type AnalyseRequest struct {
	Text string `json:"inputText"`
}

type AnalyseResponse struct {
	WordCount   int    `json:"wordCount"`
	CharCount   int    `json:"charCount"`
	LongestWord string `json:"longestWord"`
}

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