package handlers

import (
    "github.com/Jorge-Junior7/A3shopping/back-end/db"
    "github.com/Jorge-Junior7/A3shopping/back-end/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
    rows, err := db.DB.Query("SELECT id, name, price FROM products")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
        return
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var product models.Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler produtos"})
            return
        }
        products = append(products, product)
    }

    c.JSON(http.StatusOK, products)
}
