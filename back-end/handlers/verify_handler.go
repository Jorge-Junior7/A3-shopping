package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/gin-gonic/gin"
)

// Função para verificar os dados do usuário
func VerifyUserData(c *gin.Context) {
	var user models.User

	// Faz o bind dos dados recebidos em JSON e loga o erro, se houver
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Erro ao fazer o bind dos dados JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Remove pontos e traços do CPF
	re := regexp.MustCompile(`[^\d]`)
	user.CPF = re.ReplaceAllString(user.CPF, "")

	// Logs para verificar os dados tratados
	log.Printf("Email recebido: %s", user.Email)
	log.Printf("CPF (formatado) recebido: %s", user.CPF)
	log.Printf("Data de Nascimento recebida: %s", user.BirthDate)
	log.Printf("Frase de Recuperação recebida: %s", user.RecoveryPhrase)

	// Verifique se todos os campos obrigatórios estão preenchidos
	if user.Email == "" || user.CPF == "" || user.BirthDate == "" || user.RecoveryPhrase == "" {
		log.Println("Dados ausentes: verifique se email, CPF, data de nascimento e frase de recuperação estão preenchidos.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os campos são obrigatórios."})
		return
	}

	// Log dos dados recebidos para verificação
	log.Printf("Dados recebidos - Email: %s, CPF: %s, BirthDate: %s, RecoveryPhrase: %s", user.Email, user.CPF, user.BirthDate, user.RecoveryPhrase)

	// Consulta SQL para verificar se os dados existem no banco
	query := `SELECT * FROM users WHERE email = $1 AND cpf = $2 AND birthdate = $3 AND recovery_phrase = $4`
	err := db.DB.QueryRow(query, user.Email, user.CPF, user.BirthDate, user.RecoveryPhrase).Scan(
		&user.ID, &user.FullName, &user.BirthDate,
		&user.CPF, &user.Nickname, &user.Location,
		&user.Email, &user.Password, &user.RecoveryPhrase, &user.ProfilePhoto,
	)

	// Tratamento do resultado da consulta
	if err == sql.ErrNoRows {
		log.Println("Nenhum usuário encontrado com os dados fornecidos")
		c.JSON(http.StatusNotFound, gin.H{"message": "Palavra-chave não reconhecida ou não registrada"})
		return
	} else if err != nil {
		log.Printf("Erro ao executar a consulta SQL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar os dados do usuário"})
		return
	}

	log.Println("Dados verificados com sucesso")
	c.JSON(http.StatusOK, gin.H{"message": "Dados verificados com sucesso"})
}
