package main

import (
	"log"
	"net/http"

	"github.com/Jorge-Junior7/A3shopping/back-end/db"
	"github.com/Jorge-Junior7/A3shopping/back-end/routes"
)

func main() {
	// Conectar ao banco de dados
	db.Connect()

	// Definir as rotas
	router := routes.SetupRoutes()

	// Iniciar o servidor
	log.Fatal(http.ListenAndServe(":8080", router))
}
