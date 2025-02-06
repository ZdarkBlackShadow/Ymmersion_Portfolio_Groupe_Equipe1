package service

import (
	"fmt"
	"net/http"
)

func Rating(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if r.Method != http.MethodPost {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}

		// Récupérer l'ID de l'épreuve
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "ID manquant", http.StatusBadRequest)
			return
		}

		// Récupérer la note (on sait que le name sera "rating_<id>")
		rating := r.FormValue("rating_" + id)

		// Vérifier que la note est valide (1, 2 ou 3)
		if rating != "1" && rating != "2" && rating != "3" {
			http.Error(w, "Valeur de note invalide", http.StatusBadRequest)
			return
		}

		// Afficher le résultat (ici, on pourrait stocker en base de données)
		fmt.Printf("Épreuve ID: %s - Note reçue : %s\n", id, rating)

	} else {
		http.Error(w, "Méthode HTTP non autorisée", http.StatusMethodNotAllowed)
	}
}
