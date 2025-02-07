package service

import (
	"net/http"
)

// Handler pour afficher le calendrier
func calendarHandler(w http.ResponseWriter, r *http.Request) {
	schedule, err := loadEvents()
	if err != nil {
		http.Error(w, "Erreur lors du chargement des événements", http.StatusInternalServerError)
		return
	}
	err = Templates.ExecuteTemplate(w, "calendar", schedule)
	if err != nil {
		http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
		return
	}
}
