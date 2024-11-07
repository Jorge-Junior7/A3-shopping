package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// AuthMiddleware adiciona o ID do usuário ao contexto
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, err := getUserIDFromToken(c.Request)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
            return
        }
        c.Set("id", userID)
        c.Next()
    }
}

// Placeholder para a função de extração de ID (substitua pela lógica real de extração de token)
func getUserIDFromToken(r *http.Request) (string, error) {
    return "1", nil // Exemplo: retorne "1" como um ID válido para testes
}
