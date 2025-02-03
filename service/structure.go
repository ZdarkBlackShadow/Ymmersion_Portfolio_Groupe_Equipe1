package service

type Epreuve struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"image"`
	CodeGo      string `json:"codego"`
	CodeHtml    string `json:"codeHtml"`
	CodeCss     string `json:"codeCss"`
	FinalResult string `json:"finalResult"`
	Timer       int    `json:"timerInMinute"`
	Date        string `json:"when"`
}

type DataTDB struct {
	Datatdb string
}