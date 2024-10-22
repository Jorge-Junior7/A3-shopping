package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/gin-gonic/gin"
)

// AddMessage adiciona uma nova mensagem ao banco de dados
func AddMessage(c *gin.Context) {
	var msg models.Messages

	// Bind JSON input to the Messages struct
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Buscar o nickname do sender (remetente)
	var senderNickname string
	err := db.DB.QueryRow("SELECT nickname FROM users WHERE id = $1", msg.SenderID).Scan(&senderNickname)
	if err != nil {
		log.Println("Erro ao buscar o nickname do remetente:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar o nickname do remetente"})
		return
	}

	// Buscar o nickname do receiver (destinatário)
	var receiverNickname string
	err = db.DB.QueryRow("SELECT nickname FROM users WHERE id = $1", msg.ReceiverID).Scan(&receiverNickname)
	if err != nil {
		log.Println("Erro ao buscar o nickname do receptor:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar o nickname do receptor"})
		return
	}

	// Inserir a mensagem no banco de dados com os nicknames
	_, err = db.DB.Exec(
		"INSERT INTO messages (sender_id, receiver_id, message, sender_name, receiver_name) VALUES ($1, $2, $3, $4, $5)",
		msg.SenderID,
		msg.ReceiverID,
		msg.Content,
		senderNickname,   // Adicionando o nickname do remetente ao banco
		receiverNickname, // Adicionando o nickname do receptor ao banco
	)

	if err != nil {
		log.Println("Erro ao adicionar mensagem:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar mensagem"})
		return
	}

	log.Println("Mensagem adicionada com sucesso!")
	c.JSON(http.StatusOK, gin.H{"message": "Mensagem adicionada com sucesso"})
}

// GetMessages recupera todas as mensagens entre dois usuários
func GetMessages(c *gin.Context) {
	senderID := c.Param("sender_id")
	receiverID := c.Param("receiver_id")

	// Query para buscar mensagens entre sender_id e receiver_id
	rows, err := db.DB.Query(`
		SELECT m.id, m.sender_id, m.receiver_id, m.message, m.sender_name, m.receiver_name, m.sent_at
		FROM messages m 
		WHERE (m.sender_id = $1 AND m.receiver_id = $2) OR (m.sender_id = $2 AND m.receiver_id = $1)`, senderID, receiverID)

	if err != nil {
		log.Println("Erro ao buscar mensagens:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar mensagens"})
		return
	}
	defer rows.Close()

	var messages []models.Messages
	for rows.Next() {
		var msg models.Messages
		var senderName sql.NullString // Para tratar campos nulos
		var receiverName sql.NullString // Para tratar campos nulos

		// Ajustando Scan para lidar com sender_name e receiver_name nulos
		if err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &senderName, &receiverName, &msg.Timestamp); err != nil {
			log.Printf("Erro ao ler mensagens: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler mensagens"})
			return
		}

		// Se senderName não for nulo, atribuímos ao campo SenderName
		if senderName.Valid {
			msg.SenderName = senderName.String
		} else {
			msg.SenderName = ""
		}

		// Se receiverName não for nulo, atribuímos ao campo ReceiverName
		if receiverName.Valid {
			msg.ReceiverName = receiverName.String
		} else {
			msg.ReceiverName = ""
		}

		messages = append(messages, msg)
	}

	log.Println("Mensagens recuperadas com sucesso!")
	c.JSON(http.StatusOK, messages)
}
