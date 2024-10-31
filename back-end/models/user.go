package models

type User struct {
	ID              int    `json:"id"`
	FullName        string `json:"full_name"`
	BirthDate       string `json:"birth_date"`
	CPF             string `json:"cpf"`
	Nickname        string `json:"nickname"`
	Location        string `json:"location"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ProfilePhoto    string `json:"profile_photo"` // Caminho da imagem
	RecoveryPhrase  string `json:"recovery_phrase"`
}
