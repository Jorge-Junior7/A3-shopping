package routes

import (
	"github.com/Jorge-Junior7/A3shopping/back-end/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Instanciar o handler de chat
	chatHandler := handlers.NewChatHandler()

	// Definir as rotas
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/login/reset", handlers.LoginReset)

	// Rotas de chat usando o handler de chat
	router.POST("/messages", chatHandler.SendMessage) // Enviar mensagem
	router.GET("/messages/:product_id", chatHandler.GetMessages) // Recuperar mensagens por ID do produto

	// Rotas para produtos
	router.POST("/products", handlers.AddProduct)
	router.GET("/products", handlers.GetProducts)

	return router
}
