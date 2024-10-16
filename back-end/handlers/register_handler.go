package handlers

import (
    "log"
    "net/http"
    "regexp"
    "strings"

    "github.com/Jorge-Junior7/A3shopping/back-end/db"
    "github.com/Jorge-Junior7/A3shopping/back-end/models"
    "github.com/Jorge-Junior7/A3shopping/back-end/services"
    "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }

    // Objeto para armazenar erros
    errors := make(map[string]string)

    // Verificações de campos obrigatórios
    if user.FullName == "" {
        errors["full_name"] = "O nome completo é obrigatório"
    }
    if user.BirthDate == "" {
        errors["birthdate"] = "A data de nascimento é obrigatória"
    }
    if user.CPF == "" {
        errors["cpf"] = "O CPF é obrigatório"
    }
    if user.Nickname == "" {
        errors["nickname"] = "O apelido é obrigatório"
    }

    if user.Location == "" {
        errors["location"] = "A localização é obrigatória"
    }
    if user.Email == "" {
        errors["email"] = "O email é obrigatório"
    } else if !isValidEmail(user.Email) {
        errors["email"] = "Formato de email inválido"
    }
    if user.Password == "" {
        errors["password"] = "A senha é obrigatória"
    } else if len(user.Password) < 8 {
        errors["password"] = "A senha deve ter pelo menos 8 caracteres"
    }

    // Se houver erros de validação, retorná-los
    if len(errors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
        return
    }

    // Verifica se o usuário já existe
    var existingUserID int
    err := db.DB.QueryRow("SELECT id FROM users WHERE email = $1", user.Email).Scan(&existingUserID)
    if err == nil {
        errors["email"] = "Este email já está cadastrado"
        c.JSON(http.StatusConflict, gin.H{"errors": errors})
        return
    }

    // Insere o usuário no banco de dados
    if err := services.RegisterUser(user); err != nil {
        if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
            if strings.Contains(err.Error(), "users_email_key") {
                errors["email"] = "Este email já está cadastrado"
            }
            if strings.Contains(err.Error(), "users_cpf_key") {
                errors["cpf"] = "Este CPF já está cadastrado"
            }
            c.JSON(http.StatusConflict, gin.H{"errors": errors})
            return
        }
        log.Printf("Erro ao registrar usuário: %v", err) // Loga o erro detalhado
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário"})
        return
    }

    // Resposta de sucesso
    c.JSON(http.StatusOK, gin.H{"message": "Usuário registrado com sucesso!"})
}

// Função para validar o formato do email
func isValidEmail(email string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}
