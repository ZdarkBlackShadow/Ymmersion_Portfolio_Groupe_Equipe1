package service

import (
	"net/http"
)

func calendarHandler(w http.ResponseWriter, r *http.Request) {
	schedule, err := loadEvents()
	if err != nil {
		http.Error(w, "Erreur lors du chargement des événements", http.StatusInternalServerError)
		return
	}

	// Parse le fichier HTML
	err = Templates.ExecuteTemplate(w, "calendar", schedule)
	if err != nil {
		http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
		return
	}
}
