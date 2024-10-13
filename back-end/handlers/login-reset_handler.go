package handlers

import (
	"net/http"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/gin-gonic/gin"
)

// Função para redefinir a senha
func ResetPassword(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Atualizar a senha no banco de dados
	_, err := db.DB.Exec("UPDATE users SET password=$1 WHERE email=$2", user.Password, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao redefinir a senha"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Senha redefinida com sucesso!"})
}
