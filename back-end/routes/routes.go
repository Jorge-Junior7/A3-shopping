package routes

import (
	"github.com/Jorge-Junior7/A3shopping/back-end/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Definir as rotas
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/login/reset", handlers.LoginReset)

	// Rotas de mensagens
	router.POST("/messages", handlers.AddMessage) // Enviar mensagem
	router.GET("/messages/:sender_id/:receiver_id", handlers.GetMessages) // Recuperar mensagens entre dois usuários

	// Rotas para produtos
	router.POST("/products", handlers.AddProduct)
	router.GET("/products", handlers.GetProducts)

	return router
}
