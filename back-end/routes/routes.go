package routes

import (
    "github.com/Jorge-Junior7/A3shopping/back-end/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    router := gin.Default()

    // Rotas de usuário
    router.POST("/register", handlers.Register)
    router.POST("/login", handlers.Login)

    // Rotas de produto
    router.GET("/products", handlers.GetProducts)
    
    return router
}
