package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "github.com/Jorge-Junior7/A3shopping/back-end/db"
    "github.com/Jorge-Junior7/A3shopping/back-end/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Verificar se o servidor já está em execução
    if err := checkIfRunning(); err != nil {
        log.Fatal(err)
    }
    defer os.Remove("server.lock") // Remove o lock file ao encerrar

    // Definir o modo de produção (release)
    gin.SetMode(gin.ReleaseMode)

    // Conectar ao banco de dados
    db.Connect()

    // Definir as rotas
    router := routes.SetupRoutes()

    // Canal para escutar sinais de interrupção
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

    // Iniciar o servidor na porta 8080
    go func() {
        log.Fatal(http.ListenAndServe(":8080", router))
    }()

    // Esperar por um sinal de interrupção
    <-stop

    log.Println("Servidor encerrado.")
}

// checkIfRunning verifica se o servidor já está em execução
func checkIfRunning() error {
    lockFile, err := os.OpenFile("server.lock", os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666)
    if err != nil {
        return err
    }
    defer lockFile.Close()

    return nil
}
