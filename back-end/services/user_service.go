package services

import (
    "github.com/Jorge-Junior7/A3shopping/back-end/models"
    "github.com/Jorge-Junior7/A3shopping/back-end/db"
)

func RegisterUser(user models.User) error {
    _, err := db.DB.Exec("INSERT INTO users (email, password, full_name, birth_date, cpf, nickname, location, photo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", user.Email, user.Password, user.FullName, user.BirthDate, user.CPF, user.Nickname, user.Location, user.Photo)
    return err
}

// Outras funções...
