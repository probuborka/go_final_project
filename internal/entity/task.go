package entity

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type IdTask struct {
	ID string `json:"id"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}
