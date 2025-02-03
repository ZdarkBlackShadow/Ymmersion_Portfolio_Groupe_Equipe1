package service

type DataExample struct {
	Data string
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
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	CodeGo      []string `json:"codeGo"`
	CodeHTML    string   `json:"codehtml"`
	CodeCSS     string   `json:"codecss"`
	FinalImage  string   `json:"finalresult"`
}

type DataTDB struct {
	Datatdb string
}
