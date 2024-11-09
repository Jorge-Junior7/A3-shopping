package handlers

import (
	"log"
	"net/http"
	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UpdatePasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

// Função para atualizar a senha do usuário
func UpdateUserPassword(c *gin.Context) {
	var req UpdatePasswordRequest

	// Faz o bind dos dados recebidos em JSON e loga o erro, se houver
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Erro ao fazer o bind dos dados JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Criptografa a nova senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Erro ao criptografar a senha: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar a senha"})
		return
	}

	// Atualiza a senha do usuário no banco de dados com base no email
	query := `UPDATE users SET password = $1 WHERE email = $2`
	res, err := db.DB.Exec(query, string(hashedPassword), req.Email)
	if err != nil {
		log.Printf("Erro ao atualizar a senha no banco de dados: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar a senha"})
		return
	}

	// Verifica se algum registro foi atualizado
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Erro ao verificar o resultado da atualização: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar a senha"})
		return
	}
	if rowsAffected == 0 {
		log.Println("Nenhum usuário encontrado com o email fornecido")
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado"})
		return
	}

	log.Println("Senha atualizada com sucesso")
	c.JSON(http.StatusOK, gin.H{"message": "Senha atualizada com sucesso"})
}
