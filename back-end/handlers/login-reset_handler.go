package handlers

import (
	"net/http"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginReset redefine a senha do usuário
func LoginReset(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Verificar se o usuário com o email fornecido existe
	var existingUser models.User
	err := db.DB.QueryRow("SELECT id, email, recovery_phrase FROM users WHERE email=$1", user.Email).Scan(&existingUser.ID, &existingUser.Email, &existingUser.RecoveryPhrase)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// Verificar se a frase de recuperação está correta
	if user.RecoveryPhrase != existingUser.RecoveryPhrase {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Frase de recuperação incorreta"})
		return
	}

	// Verificar a força da nova senha
	if !isValidPassword(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A nova senha deve ter pelo menos 8 caracteres, uma letra maiúscula, uma letra minúscula, um número e um caractere especial."})
		return
	}

	// Hash da nova senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}

	// Atualizar a senha no banco de dados
	_, err = db.DB.Exec("UPDATE users SET password=$1 WHERE email=$2", hashedPassword, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao redefinir a senha"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Senha redefinida com sucesso!"})
}
