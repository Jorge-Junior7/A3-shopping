package models

type Product struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Category    string  `json:"category"` 
    Location    string  `json:"location"`
    Photo       string  `json:"photo"`
}
