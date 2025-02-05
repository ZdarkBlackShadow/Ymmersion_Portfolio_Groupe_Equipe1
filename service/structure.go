package service

//Structures pour les membres du groupe
type User struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Id                 int    `json:"id"`
	Img                string `json:"img"`
	RoleDescription    string `json:"roleDescription"`
	EpreuveName        string `json:"epreuveName"`
	DescriptionEpreuve string `json:"descriptionEpreuve"`
	Epreuve            string `json:"epreuve"`
}
type UserStruct struct {
	List []string `json:"list"`
	Data []User   `json:"data"`
}

//Structure pour les epreuves
type Exercice struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Image            string   `json:"image"`
	ExempleRenduCode string   `json:"ExempleRendu"`
	CodeGo           string   `json:"codeGo"`
	CodeHTML         string   `json:"codehtml"`
	CodeCSS          string   `json:"codecss"`
	FinalImage       string   `json:"finalresult"`
	FinalResultCode  string   `json:"finalResultCode"`
	Timer            int      `json:"timerInMinute"`
	When             string   `json:"when"`
	LearningLesson   []string `json:"learningLesson"`
}

type DataTDB struct {
	Datatdb string
}

//Structure pour le calendrier

// Structure pour représenter un événement
type Event struct {
	Heure       string `json:"heure"`
	Titre       string `json:"titre"`
	Description string `json:"description"`
}

// Structure pour représenter les événements de chaque jour
type WeekSchedule struct {
	Lundi    []Event `json:"lundi"`
	Mardi    []Event `json:"mardi"`
	Mercredi []Event `json:"mercredi"`
	Jeudi    []Event `json:"jeudi"`
	Vendredi []Event `json:"vendredi"`
}
