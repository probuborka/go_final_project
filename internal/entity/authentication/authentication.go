package entityauth

type Authentication struct {
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}
