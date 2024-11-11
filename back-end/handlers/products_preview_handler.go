package handlers

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "path/filepath"
    "strings"

    "github.com/Jorge-Junior7/A3shopping/back-end/db"
    "github.com/gin-gonic/gin"
)

const imageBaseURL = "http://localhost:8080/uploads_products"

// GetProductsPreview busca campos photo1, title, publish_date e price dos produtos no banco de dados e retorna como JSON
func GetProductsPreview(c *gin.Context) {
    rows, err := db.DB.Query(`
        SELECT title, COALESCE(publish_date::text, '') AS publish_date, photo1, price 
        FROM products
    `)
    if err != nil {
        log.Printf("Erro ao buscar produtos: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
        return
    }
    defer rows.Close()

    var products []gin.H

    for rows.Next() {
        var title string
        var publishDate string
        var photo1 sql.NullString
        var price sql.NullFloat64

        err := rows.Scan(&title, &publishDate, &photo1, &price)
        if err != nil {
            log.Printf("Erro ao ler dados do produto: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler dados do produto"})
            return
        }

        // Monta o caminho completo da imagem principal `photo1`, se disponível
        photo1Path := ""
        if photo1.Valid && photo1.String != "" {
            cleanPhoto1 := strings.Trim(photo1.String, "<>")
            photo1Path = fmt.Sprintf("%s/%s", imageBaseURL, filepath.Base(cleanPhoto1))
        } else {
            log.Printf("Imagem principal ausente ou inválida para o produto: %s", title)
            // Opcionalmente, você pode definir uma imagem padrão aqui
        }

        products = append(products, gin.H{
            "title":        title,
            "publish_date": publishDate,
            "photo1":       photo1Path,
            "price":        price.Float64,
        })
    }

    c.JSON(http.StatusOK, products)
}
