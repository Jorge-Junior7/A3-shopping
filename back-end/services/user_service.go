package services

import (
	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
)

// RegisterUser insere um novo usuário no banco de dados
func RegisterUser(user models.User) error {
	query := `INSERT INTO users (full_name, birthdate, cpf, nickname, location, email, password, recovery_phrase) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := db.DB.Exec(query, user.FullName, user.BirthDate, user.CPF, user.Nickname, user.Location, user.Email, user.Password, user.RecoveryPhrase)
	return err
}
