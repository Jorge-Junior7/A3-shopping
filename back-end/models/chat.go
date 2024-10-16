package models

type Message struct {
    ID          int    `json:"id"`
    ProductID   int    `json:"product_id"`
    SenderID    int    `json:"sender_id"`
    ReceiverID  int    `json:"receiver_id"`
    Content     string `json:"content"`
    Timestamp   string `json:"timestamp"` // ou time.Time, se você preferir trabalhar com tipos de data
}
