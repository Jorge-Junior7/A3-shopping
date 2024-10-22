// Nome do arquivo: chat_model.go
package models

// Message representa uma mensagem de chat
type Message struct {
	ID          int    `json:"id"`
	SenderID    int    `json:"sender_id"`
	ReceiverID  int    `json:"receiver_id"`
	Content     string `json:"content"`
	Timestamp   string `json:"timestamp"` // ou time.Time, se preferir trabalhar com tipos de data
}
