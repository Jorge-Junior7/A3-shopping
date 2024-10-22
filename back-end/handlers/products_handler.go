package handlers

import (
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

	// Validação de campos obrigatórios
	var missingFields []string

	if product.Title == "" {
		missingFields = append(missingFields, "título")
	}
	if product.Description == "" {
		missingFields = append(missingFields, "descrição")
	}
	if product.Price == 0 {
		missingFields = append(missingFields, "preço")
	}
	if product.Category == "" {
		missingFields = append(missingFields, "categoria")
	}
	if product.Location == "" {
		missingFields = append(missingFields, "localização")
	}

	// Foto não é obrigatória, então não a validamos

	// Verifica se há campos faltando
	if len(missingFields) > 0 {
		// Se houver mais de um campo faltando
		if len(missingFields) > 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Preencha os campos obrigatórios:",
			})
		} else {
			// Se houver apenas um campo faltando
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Preencha o campo " + missingFields[0],
			})
		}
		return
	}

	// Tenta inserir o produto no banco de dados
	query := `INSERT INTO products (title, description, price, category, location, photo) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.DB.QueryRow(query,
		product.Title,
		product.Description,
		product.Price,
		product.Category,
		product.Location,
		product.Photo,
	).Scan(&product.ID)

	// Verifica se houve erro na inserção
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar o produto", "details": err.Error()})
		return
	}

	// Resposta de sucesso
	c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado com sucesso", "product_id": product.ID})
}

// GetProducts recupera todos os produtos do banco de dados.
func GetProducts(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, description, price, category, location, photo FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Category, &product.Location, &product.Photo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler produtos"})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}
