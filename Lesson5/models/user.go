package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
