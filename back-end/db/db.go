package db

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"         // Importa o driver PostgreSQL
    "github.com/joho/godotenv"    // Importa a biblioteca para ler o arquivo .env
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
    
    // A string de conexão é geralmente construída com as variáveis
    connStr := os.Getenv("DB_CONN_STR")
    if connStr == "" {
        // Se a string de conexão não estiver definida no .env, você pode construí-la a partir das variáveis
        connStr = "user=" + os.Getenv("DB_USER") +
            " password=" + os.Getenv("DB_PASSWORD") +
            " dbname=" + os.Getenv("DB_NAME") +
            " host=" + os.Getenv("DB_HOST") +
            " port=" + os.Getenv("DB_PORT") +
            " sslmode=disable"
    }

    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Erro ao conectar ao banco de dados: ", err)
    }

    // Verifica se a conexão é válida
    if err = DB.Ping(); err != nil {
        log.Fatal("Banco de dados não está acessível: ", err)
    }

    log.Println("Conectado ao banco de dados com sucesso!")
}
