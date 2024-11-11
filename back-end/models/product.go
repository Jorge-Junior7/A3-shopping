package models

import "database/sql"

type Product struct {
    ID          int            `json:"id"`
    Title       string         `json:"title"`
    Description string         `json:"description"`
    Price       float64        `json:"price"`
    Category    string         `json:"category"`
    Condition   sql.NullString `json:"condition"`  // Alterado para sql.NullString
    UserID      int            `json:"user_id"`
    Photo1      sql.NullString `json:"photo1"`
    Photo2      sql.NullString `json:"photo2"`
    Photo3      sql.NullString `json:"photo3"`
    Photo4      sql.NullString `json:"photo4"`
}
