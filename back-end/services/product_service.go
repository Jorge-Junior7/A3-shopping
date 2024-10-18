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

	// Bind JSON input to the Product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Verifique se o user_id e o category_id existem
	if !validateForeignKey("users", product.UserID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id inválido"})
		return
	}
	if !validateForeignKey("categories", product.CategoryID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category_id inválido"})
		return
	}

	// Inserir o produto no banco de dados
	query := `
		INSERT INTO products (title, description, price, category_id, user_id, location, photo)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`

	err := db.DB.QueryRow(query,
		product.Title,
		product.Description,
		product.Price,
		product.CategoryID,
		product.UserID,
		product.Location,
		product.Photo,
	).Scan(&product.ID)

	if err != nil {
		log.Printf("Erro ao adicionar produto: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar o produto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado com sucesso", "product_id": product.ID})
}

// GetProducts recupera todos os produtos do banco de dados.
func GetProducts(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, description, price, category_id, user_id, location, photo FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.CategoryID, &product.UserID, &product.Location, &product.Photo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler produtos"})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

// validateForeignKey verifica se a chave estrangeira existe no banco de dados.
func validateForeignKey(table string, id int) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM ` + table + ` WHERE id = $1)`
	err := db.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		log.Printf("Erro ao verificar chave estrangeira: %v", err)
		return false
	}
	return exists
}
