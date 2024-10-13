package services

import (
    "github.com/Jorge-Junior7/A3shopping/back-end/models"
    "github.com/Jorge-Junior7/A3shopping/back-end/db"
)

func RegisterUser(user models.User) error {
    // Ajuste os nomes das colunas para corresponder ao banco de dados
    _, err := db.DB.Exec(
        "INSERT INTO users (full_name, birthdate, cpf, nickname, profile_photo, location, email, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
        user.FullName, 
        user.BirthDate, 
        user.CPF, 
        user.Nickname, 
        user.ProfilePhoto, 
        user.Location, 
        user.Email, 
        user.Password,
    )
    
    if err != nil {
        return err // Retorna o erro para tratamento posterior
    }

    return nil // Retorna nil se a inserção for bem-sucedida
}
