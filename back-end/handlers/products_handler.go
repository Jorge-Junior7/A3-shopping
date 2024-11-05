// handlers/product_handler.go
package handlers

import (
    "database/sql"
    "net/http"
    "strconv"
    "strings"

    "github.com/Jorge-Junior7/A3shopping/back-end/db"
    "github.com/Jorge-Junior7/A3shopping/back-end/models"
    "github.com/gin-gonic/gin"
)

// AddProduct adiciona um novo produto ao banco de dados.
func AddProduct(c *gin.Context) {
    var product models.Product

    // Faz o Bind do JSON para a estrutura de produto
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
        return
    }

    // Obtém o user_id a partir do contexto (middleware de autenticação)
    userIDStr, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
        return
    }

    userID, err := strconv.Atoi(userIDStr.(string))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
        return
    }
    product.UserID = userID

    // Validação de campos obrigatórios
    var missingFields []string
    if product.Title == "" { missingFields = append(missingFields, "título") }
    if product.Description == "" { missingFields = append(missingFields, "descrição") }
    if product.Price == 0 { missingFields = append(missingFields, "preço") }
    if product.Category == "" { missingFields = append(missingFields, "categoria") }
    if product.Condition == "" { missingFields = append(missingFields, "condição") }
    
    if len(missingFields) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Preencha os campos obrigatórios: " + strings.Join(missingFields, ", "),
        })
        return
    }

    // Insere o produto no banco de dados
    query := `
        INSERT INTO products (title, description, price, category, condition, user_id, photo1, photo2, photo3, photo4) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id
    `
    err = db.DB.QueryRow(query,
        product.Title, product.Description, product.Price, product.Category, product.Condition, product.UserID,
        sql.NullString{String: *product.Photo1, Valid: product.Photo1 != nil},
        sql.NullString{String: *product.Photo2, Valid: product.Photo2 != nil},
        sql.NullString{String: *product.Photo3, Valid: product.Photo3 != nil},
        sql.NullString{String: *product.Photo4, Valid: product.Photo4 != nil},
    ).Scan(&product.ID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar o produto", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado com sucesso", "product_id": product.ID})
}

// GetProducts recupera todos os produtos do banco de dados.
func GetProducts(c *gin.Context) {
    rows, err := db.DB.Query(`
        SELECT id, title, description, price, category, condition, user_id, photo1, photo2, photo3, photo4 
        FROM products
    `)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos", "details": err.Error()})
        return
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var product models.Product
        var photo1, photo2, photo3, photo4 sql.NullString

        if err := rows.Scan(
            &product.ID, &product.Title, &product.Description, &product.Price,
            &product.Category, &product.Condition, &product.UserID,
            &photo1, &photo2, &photo3, &photo4,
        ); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler produtos", "details": err.Error()})
            return
        }

        // Define as fotos para nil se estiverem nulas no banco de dados
        if photo1.Valid { product.Photo1 = &photo1.String }
        if photo2.Valid { product.Photo2 = &photo2.String }
        if photo3.Valid { product.Photo3 = &photo3.String }
        if photo4.Valid { product.Photo4 = &photo4.String }

        products = append(products, product)
    }

    c.JSON(http.StatusOK, products)
}
