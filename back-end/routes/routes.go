package routes

import (
    "github.com/Jorge-Junior7/A3shopping/back-end/handlers"
    "github.com/Jorge-Junior7/A3shopping/back-end/middleware"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:4200"},
        AllowMethods:     []string{"POST", "GET", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    router.Use(middleware.AuthMiddleware())

    router.POST("/register", handlers.Register)
    router.POST("/login", handlers.Login)
    router.POST("/login/reset", handlers.LoginReset)

    router.POST("/messages", handlers.AddMessage)
    router.GET("/messages/:sender_id/:receiver_id", handlers.GetMessages)

    router.POST("/products", handlers.AddProduct)
    router.GET("/products", handlers.GetProducts)

    router.GET("/products/preview", handlers.GetProductsPreview)

    // Rota estática para servir as imagens de produtos na nova localização
    router.Static("/uploads_products", "./handlers/uploads_products")

    router.POST("/register/verify", handlers.VerifyUserData)
    router.POST("/register/update-password", handlers.UpdateUserPassword)

    return router
}
