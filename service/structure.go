package service

type Epreuve struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Img             string `json:"image"`
	CodeGo          string `json:"codego"`
	CodeHtml        string `json:"codeHtml"`
	CodeCss         string `json:"codeCss"`
	FinalResult     string `json:"finalResult"`
	FinalResultCode string `json:"finalResultCode"`
	Timer           int    `json:"timerInMinute"`
	Date            string `json:"when"`
}

type User struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Id              int    `json:"id"`
	RoleDescription string `json:"roleDescription"`
}
type UserStruct struct {
	List []string `json:"list"`
	Data []User   `json:"data"`
}

type Exercice struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Image            string `json:"image"`
	ExempleRenduCode string `json:"ExempleRendu"`
	CodeGo           string `json:"codeGo"`
	CodeHTML         string `json:"codehtml"`
	CodeCSS          string `json:"codecss"`
	FinalImage       string `json:"finalresult"`
	FinalResultCode  string `json:"finalResultCode"`
	Timer            int    `json:"timerInMinute"`
	When             string `json:"when"`
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
