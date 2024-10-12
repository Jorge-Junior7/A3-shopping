package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "os"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
    // Carrega o arquivo .env do caminho específico
    err := godotenv.Load("/home/junior/Documentos/A3shopping/back-end/.env")
    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env: ", err)
    }

    // Log das variáveis de ambiente
    log.Println("DB_HOST:", os.Getenv("DB_HOST"))
    log.Println("DB_PORT:", os.Getenv("DB_PORT"))
    log.Println("DB_USER:", os.Getenv("DB_USER"))
    log.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
    log.Println("DB_NAME:", os.Getenv("DB_NAME"))
    log.Println("DB_CONN_STR:", os.Getenv("DB_CONN_STR"))

    // Pega a string de conexão do .env
    connStr := os.Getenv("DB_CONN_STR")
    if connStr == "" {
        log.Fatal("String de conexão não encontrada. Verifique o arquivo .env.")
    }

    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Erro ao conectar ao banco de dados: ", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Banco de dados não está acessível: ", err)
    }

    log.Println("Conectado ao banco de dados com sucesso!")
}
