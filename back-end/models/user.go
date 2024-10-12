package models

type User struct {
    ID            int     `db:"id"`
    Email         string  `db:"email"`
    Password      string  `db:"password"`
    FullName      string  `db:"full_name"`      // Nome completo
    BirthDate     string  `db:"birth_date"`     // Data de nascimento
    CPF           string  `db:"cpf"`            // CPF
    Nickname      string  `db:"nickname"`       // Apelido
    Location      string  `db:"location"`       // Localização
    Photo         string  `db:"photo"`          // Foto
}