package main

import (
    "github.com/Jorge-Junior7/A3shopping/back-end/db" // Corrigido para refletir o novo nome do módulo
    "log"
)

func main() {
    err := db.Connect()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
}
