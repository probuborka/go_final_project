package entity

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type ID struct {
	ID string `json:"id"`
}

type Error struct {
	Error string `json:"error"`
}
