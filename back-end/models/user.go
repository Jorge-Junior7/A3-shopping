package models

type User struct {
	ID           int    `json:"id"`
	FullName     string `json:"full_name"`
	BirthDate    string `json:"birthdate"`
	CPF          string `json:"cpf"`
	Nickname     string `json:"nickname"`
	ProfilePhoto string `json:"photo"`
	Location     string `json:"location"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}
