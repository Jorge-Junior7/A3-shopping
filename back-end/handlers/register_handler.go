package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/models"
	"github.com/Jorge-Junior7/A3shopping/back-end/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Função para lidar com o upload de imagens
func UploadImage(c *gin.Context) (string, error) {
	file, err := c.FormFile("profilePhoto")
	if err != nil {
		log.Printf("Erro ao obter o arquivo: %v", err)
		return "", fmt.Errorf("erro ao obter o arquivo: %w", err)
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		log.Printf("Tipo de arquivo inválido: %s", ext)
		return "", fmt.Errorf("tipo de arquivo inválido; apenas JPG e PNG são aceitos")
	}

	deviceName := strings.Split(file.Filename, ".")[0]
	newFileName := fmt.Sprintf("%s_%d%s", deviceName, time.Now().UnixNano(), ext)
	dst := filepath.Join("uploads", newFileName)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Printf("Erro ao salvar o arquivo: %v", err)
		return "", fmt.Errorf("erro ao salvar o arquivo: %w", err)
	}

	log.Printf("Arquivo salvo com sucesso: %s", dst)
	return dst, nil
}

// Função para gerar uma frase de recuperação aleatória
func generateRecoveryPhrase() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		log.Println("Erro ao gerar frase de recuperação:", err)
		return ""
	}
	return hex.EncodeToString(bytes)
}

// Função de Registro do Usuário
func Register(c *gin.Context) {
	// Processar os dados do formulário multipart
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		log.Printf("Erro ao fazer parsing do formulário: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Verifique e obtenha cada campo do formulário
	user := models.User{
		FullName:  c.PostForm("full_name"),
		BirthDate: c.PostForm("birthdate"),
		CPF:       c.PostForm("cpf"),
		Nickname:  c.PostForm("nickname"),
		Location:  c.PostForm("location"),
		Email:     c.PostForm("email"),
		Password:  c.PostForm("password"),
	}

	// Logs para verificar cada campo recebido
	log.Printf("Nome completo: %s", user.FullName)
	log.Printf("Data de nascimento: %s", user.BirthDate)
	log.Printf("CPF: %s", user.CPF)
	log.Printf("Apelido: %s", user.Nickname)
	log.Printf("Localização: %s", user.Location)
	log.Printf("Email: %s", user.Email)
	log.Printf("Senha: %s", user.Password)

	fileName, err := UploadImage(c)
	if err != nil {
		log.Printf("Erro ao fazer upload da imagem: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ProfilePhoto = fileName

	errors := make(map[string]string)

	// Verificação de campos obrigatórios e validações específicas
	if user.FullName == "" {
		errors["full_name"] = "O nome completo é obrigatório"
	}
	if user.BirthDate == "" {
		errors["birthdate"] = "A data de nascimento é obrigatória"
	}
	if user.CPF == "" {
		errors["cpf"] = "O CPF é obrigatório"
	} else if !isValidCPF(user.CPF) {
		errors["cpf"] = "CPF inválido"
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
	} else if !isValidPassword(user.Password) {
		errors["password"] = "A senha deve ter pelo menos 8 caracteres, uma letra maiúscula, uma letra minúscula, um número e um caractere especial."
	}

	if len(errors) > 0 {
		log.Printf("Erros de validação: %+v", errors)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	var existingUserID int
	err = db.DB.QueryRow("SELECT id FROM users WHERE email = $1", user.Email).Scan(&existingUserID)
	if err == nil {
		errors["email"] = "Este email já está cadastrado"
		c.JSON(http.StatusConflict, gin.H{"errors": errors})
		return
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar se o usuário existe"})
		return
	}

	user.RecoveryPhrase = generateRecoveryPhrase()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário"})
		return
	}
	user.Password = string(hashedPassword)

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Usuário registrado com sucesso!",
		"recovery_phrase": user.RecoveryPhrase,
	})
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func isValidCPF(cpf string) bool {
	re := regexp.MustCompile(`[^\d]`)
	cpf = re.ReplaceAllString(cpf, "")
	return len(cpf) == 11
}
