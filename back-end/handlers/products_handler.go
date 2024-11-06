package handlers

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"

    "github.com/Jorge-Junior7/A3shopping/back-end/db"
    "github.com/Jorge-Junior7/A3shopping/back-end/models"
    "github.com/gin-gonic/gin"
)

// Lista de categorias e condições válidas
var validCategories = []string{"Veículos", "Locações", "Roupas e Calçados", "Móveis", "Eletrônicos", "Outros"}
var validConditions = []string{"Novo", "Semi-novo", "Em Bom Estado", "Condições Razoáveis", "Extensivo"}

// Função para verificar se um valor está na lista de valores válidos
func isValid(value string, validValues []string) bool {
    value = strings.TrimSpace(strings.ToLower(value))
    for _, v := range validValues {
        if strings.ToLower(v) == value {
            return true
        }
    }
    return false
}

// Função para salvar uma imagem e retornar o nome do arquivo
func UploadImages(c *gin.Context, fileKey string) (string, error) {
    file, err := c.FormFile(fileKey)
    if err != nil {
        fmt.Printf("Erro ao obter o arquivo %s: %v\n", fileKey, err)
        return "", nil // Continua sem erro se o arquivo não for enviado
    }

    fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
    filePath := filepath.Join("uploads_products", fileName)

    if _, err := os.Stat("uploads_products"); os.IsNotExist(err) {
        if mkErr := os.MkdirAll("uploads_products", os.ModePerm); mkErr != nil {
            return "", fmt.Errorf("erro ao criar o diretório: %w", mkErr)
        }
    }

    if err := c.SaveUploadedFile(file, filePath); err != nil {
        fmt.Printf("Erro ao salvar o arquivo %s: %v\n", fileKey, err)
        return "", fmt.Errorf("erro ao salvar o arquivo: %w", err)
    }

    return fileName, nil
}

func AddProduct(c *gin.Context) {
    fmt.Println("Iniciando AddProduct...")

    title := c.PostForm("title")
    description := c.PostForm("description")
    priceStr := c.PostForm("price")
    category := c.PostForm("category")
    condition := c.PostForm("condition")

    fmt.Println("Valores recebidos:")
    fmt.Println("Título:", title)
    fmt.Println("Descrição:", description)
    fmt.Println("Preço:", priceStr)
    fmt.Println("Categoria:", category)
    fmt.Println("Condição:", condition)

    price, err := strconv.ParseFloat(priceStr, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Preço inválido", "details": err.Error()})
        return
    }

    if !isValid(category, validCategories) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Categoria inválida"})
        return
    }
    if condition != "" && !isValid(condition, validConditions) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Condição inválida"})
        return
    }

    var missingFields []string
    if title == "" {
        missingFields = append(missingFields, "título")
    }
    if price == 0 {
        missingFields = append(missingFields, "preço")
    }

    if len(missingFields) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Preencha os campos obrigatórios: " + strings.Join(missingFields, ", ")})
        return
    }

    userIDStr, exists := c.Get("id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
        return
    }

    userID, err := strconv.Atoi(userIDStr.(string))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
        return
    }

    photo1, _ := UploadImages(c, "photo1")
    photo2, _ := UploadImages(c, "photo2")
    photo3, _ := UploadImages(c, "photo3")
    photo4, _ := UploadImages(c, "photo4")

    query := `
        INSERT INTO products (title, description, price, category, condition, user_id, photo1, photo2, photo3, photo4) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id
    `
    var productID int
    err = db.DB.QueryRow(query, title, description, price, category, condition, userID, photo1, photo2, photo3, photo4).Scan(&productID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar o produto", "details": err.Error()})
        return
    }

    fmt.Println("Produto adicionado com sucesso:", productID)
    c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado com sucesso", "product_id": productID})
}

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
        err := rows.Scan(
            &product.ID, &product.Title, &product.Description, &product.Price,
            &product.Category, &product.Condition, &product.UserID,
            &product.Photo1, &product.Photo2, &product.Photo3, &product.Photo4,
        )
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler produto", "details": err.Error()})
            return
        }
        products = append(products, product)
    }

    c.JSON(http.StatusOK, products)
}
