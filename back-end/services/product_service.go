package services

import (
	"log"
	"net/http"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/gin-gonic/gin"
)

// AddProduct adiciona um novo produto ao banco de dados.
func AddProduct(c *gin.Context) {
	var product models.Product

	// Faz o Bind do JSON para a estrutura de produto
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Tenta inserir o produto no banco de dados
	query := `
		INSERT INTO products (title, description, price, category, location, photo)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`

	// Executa a query e obtém o id gerado
	err := db.DB.QueryRow(query,
		product.Title,
		product.Description,
		product.Price,
		product.Category,
		product.Location,
		product.Photo,
	).Scan(&product.ID)

	if err != nil {
		log.Printf("Erro ao adicionar produto: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar o produto", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado com sucesso", "product_id": product.ID})
}
