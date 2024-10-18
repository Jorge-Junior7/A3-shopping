package handlers

import (
	"net/http"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var product models.Product

	// Faz o Bind do JSON para a estrutura de produto
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Tenta inserir o produto no banco de dados
	query := `INSERT INTO products (title, description, price, category_id, user_id, location, photo) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.DB.Exec(query, 
		product.Title,
		product.Description,
		product.Price,
		product.CategoryID,
		product.UserID,
		product.Location,
		product.Photo,
	)

	// Verifica se houve erro na inserção
	if err != nil {
		// Logar o erro para rastreamento
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar o produto", "details": err.Error()})
		return
	}

	// Resposta de sucesso
	c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado com sucesso"})
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
