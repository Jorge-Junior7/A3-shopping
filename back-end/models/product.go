package models

type Product struct {
    ID          int      `json:"id"`
    Title       string   `json:"title" binding:"required"`
    Description string   `json:"description"`
    Price       float64  `json:"price" binding:"required"`
    Category    string   `json:"category"`
    Condition   string   `json:"condition"`
    UserID      int      `json:"user_id"`
    Photo1      *string  `json:"photo1"`
    Photo2      *string  `json:"photo2"`
    Photo3      *string  `json:"photo3"`
    Photo4      *string  `json:"photo4"`
}