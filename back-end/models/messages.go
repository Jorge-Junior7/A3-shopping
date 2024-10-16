package models

// Message representa uma mensagem de chat
type Messages struct {
	ID          int    `json:"id"`
	ProductID   int    `json:"product_id"`
	SenderID    int    `json:"sender_id"`
	ReceiverID  int    `json:"receiver_id"`
	Content     string `json:"content"`
}
