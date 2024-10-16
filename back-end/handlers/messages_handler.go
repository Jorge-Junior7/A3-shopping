package handlers

import (
	"net/http"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/gin-gonic/gin"
)

// AddMessage adiciona uma nova mensagem ao banco de dados
func AddMessage(c *gin.Context) {
	var msg models.Message

	// Bind JSON input to the Message struct
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Inserir a mensagem no banco de dados
	_, err := db.DB.Exec(
		"INSERT INTO messages (product_id, sender_id, receiver_id, content) VALUES ($1, $2, $3, $4)",
		msg.ProductID,
		msg.SenderID,
		msg.ReceiverID,
		msg.Content,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar mensagem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mensagem adicionada com sucesso"})
}

// GetMessages recupera todas as mensagens relacionadas a um produto específico
func GetMessages(c *gin.Context) {
	productID := c.Param("product_id")

	rows, err := db.DB.Query("SELECT id, product_id, sender_id, receiver_id, content FROM messages WHERE product_id = $1", productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar mensagens"})
		return
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.ProductID, &msg.SenderID, &msg.ReceiverID, &msg.Content); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler mensagens"})
			return
		}
		messages = append(messages, msg)
	}

	c.JSON(http.StatusOK, messages)
}
