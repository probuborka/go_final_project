package entity

type Authorization struct {
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}
