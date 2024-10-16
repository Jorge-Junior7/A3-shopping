package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// Message representa uma mensagem de chat
type Message struct {
	ProductID   int    `json:"product_id"`
	SenderID    int    `json:"sender_id"`
	ReceiverID  int    `json:"receiver_id"`
	Content     string `json:"content"`
}

// ChatHandler gerencia as mensagens de chat
type ChatHandler struct {
	messages []Message
	mu       sync.Mutex
}

// NewChatHandler cria um novo handler de chat
func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		messages: make([]Message, 0),
	}
}

// SendMessage processa o envio de uma mensagem de chat
func (ch *ChatHandler) SendMessage(c *gin.Context) {
	var msg Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	ch.mu.Lock()
	defer ch.mu.Unlock()

	// Adiciona a mensagem à lista
	ch.messages = append(ch.messages, msg)

	c.JSON(http.StatusOK, gin.H{"message": "Mensagem enviada com sucesso"})
}

// GetMessages recupera todas as mensagens relacionadas a um produto específico
func (ch *ChatHandler) GetMessages(c *gin.Context) {
	productIDStr := c.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto inválido"})
		return
	}

	ch.mu.Lock()
	defer ch.mu.Unlock()

	// Filtra mensagens pelo ID do produto
	var filteredMessages []Message
	for _, msg := range ch.messages {
		if msg.ProductID == productID {
			filteredMessages = append(filteredMessages, msg)
		}
	}

	c.JSON(http.StatusOK, filteredMessages)
}
