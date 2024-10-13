package handlers

import (
	"net/http"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	var storedUser models.User
	err := db.DB.QueryRow("SELECT id, email, password FROM users WHERE email=$1", user.Email).Scan(&storedUser.ID, &storedUser.Email, &storedUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha incorretos"})
		return
	}

	if storedUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha incorretos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login bem-sucedido"})
}
