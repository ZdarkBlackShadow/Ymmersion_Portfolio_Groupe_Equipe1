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

type EpreuveStruct struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	LogiqueAndReflexion string
	Resultat            string
	LearningLesson      []string
}

type User struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          int    `json:"id"`
}
type UserStruct struct {
	List []string `json:"list"`
	Data []User   `json:"data"`
}

type Exercice struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Image           string `json:"image"`
	CodeGo          string `json:"codeGo"`
	CodeHTML        string `json:"codehtml"`
	CodeCSS         string `json:"codecss"`
	FinalImage      string `json:"finalresult"`
	FinalResultCode string `json:"finalResultCode"`
	Timer           int    `json:"timerInMinute"`
	When            string `json:"when"`
}

type DataTDB struct {
	Datatdb string
}

