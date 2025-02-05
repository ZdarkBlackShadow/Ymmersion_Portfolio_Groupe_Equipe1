package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type RatingData struct {
	Epreuve1 int `json:"epreuve1"`
	Epreuve2 int `json:"epreuve2"`
	Epreuve3 int `json:"epreuve3"`
	Epreuve4 int `json:"epreuve4"`
	Epreuve5 int `json:"epreuve5"`
	Epreuve6 int `json:"epreuve6"`
}

func Traitement() {
	http.HandleFunc("/traitement", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			ep1 := r.FormValue("rating1") // Valeur pour la première épreuve
			ep2 := r.FormValue("rating2") // Valeur pour la deuxième épreuve
			ep3 := r.FormValue("rating3") // Valeur pour la troisième épreuve
			ep4 := r.FormValue("rating4") // Valeur pour la quatrième épreuve
			ep5 := r.FormValue("rating5") // Valeur pour la cinquième épreuve
			ep6 := r.FormValue("rating6") // Valeur pour la sixième épreuve

			var data RatingData
			fmt.Sscanf(ep1, "%d", &data.Epreuve1)
			fmt.Sscanf(ep2, "%d", &data.Epreuve2)
			fmt.Sscanf(ep3, "%d", &data.Epreuve3)
			fmt.Sscanf(ep4, "%d", &data.Epreuve4)
			fmt.Sscanf(ep5, "%d", &data.Epreuve5)
			fmt.Sscanf(ep6, "%d", &data.Epreuve6)

			file, err := os.OpenFile("notation.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				http.Error(w, "Erreur d'ouverture de fichier", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(data); err != nil {
				http.Error(w, "Erreur d'écriture dans le fichier", http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Les évaluations ont été sauvegardées.")
		} else {
			http.Error(w, "Méthode HTTP non autorisée", http.StatusMethodNotAllowed)
		}
	})
}
