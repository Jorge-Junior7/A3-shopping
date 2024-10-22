package models

import (
	"time"
)

type Messages struct {
	ID           int       `json:"id"`
	SenderID     int       `json:"sender_id"`
	ReceiverID   int       `json:"receiver_id"`
	Content      string    `json:"content"`
	SenderName   string    `json:"sender_name"`
	ReceiverName string    `json:"receiver_name"` // Adiciona o campo ReceiverName
	Timestamp    time.Time `json:"timestamp"`
}
