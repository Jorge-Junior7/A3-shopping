package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Message representa uma mensagem de chat
type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ch.mu.Lock()
	ch.messages = append(ch.messages, msg)
	ch.mu.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

// GetMessages retorna as mensagens de chat
func (ch *ChatHandler) GetMessages(c *gin.Context) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	c.JSON(http.StatusOK, ch.messages)
}
