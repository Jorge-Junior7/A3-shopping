package handlers

import (
	"log"
	"net/http"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Função de Login do Usuário
func Login(c *gin.Context) {
	var input models.LoginInput

	// Bind JSON para o modelo de entrada
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	var user models.User

	// Busca o usuário pelo e-mail
	err := db.DB.QueryRow("SELECT id, full_name, password FROM users WHERE email = $1", input.Email).Scan(&user.ID, &user.FullName, &user.Password)
	if err != nil {
		log.Println("Erro ao buscar usuário:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha incorretos"})
		return
	}

	// Compara a senha fornecida com a senha armazenada
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha incorretos"})
		return
	}

	// Login bem-sucedido
	c.JSON(http.StatusOK, gin.H{"message": "Login realizado com sucesso"})
}
