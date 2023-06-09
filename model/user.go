package model

type User struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	IsActive bool   `json:"-"`
	Password string `json:"password"`
}
