package models

type Product struct {
    ID          int     `db:"id"`
    Name        string  `db:"name"`
    Description string  `db:"description"` // Descrição do produto
    Price       float64 `db:"price"`
    Photos      string  `db:"photos"` // URLs ou caminhos das fotos
    CreatedAt   string  `db:"created_at"` // Data e hora de publicação
}
