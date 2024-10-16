package models

type Product struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    CategoryID  int     `json:"category_id"`
    UserID      int     `json:"user_id"`
    Location    string  `json:"location"`
    Photo       string  `json:"photo"`
}
