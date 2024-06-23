package main

import (
	"fmt"
	"log"
	"net/http"
	"ps-backend-felipe-rodrigues/src/database"
	"ps-backend-felipe-rodrigues/src/router"
)

//TODO: IMPLEMENTAR O MONGODB NESSE CÓDIGO

func main() {
	fmt.Println("Server is starting...")

	// String de conexão do MongoDB
	uri := "mongodb+srv://felipe:1234@furz-parfum.jwzwsza.mongodb.net/mydatabase?retryWrites=true&w=majority&ssl=true"

	// Conectar ao MongoDB
	_, err := database.ConectarMongo(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Gerar as rotas do servidor
	r := router.Gerar()

	// Iniciar o servidor HTTP
	log.Fatal(http.ListenAndServe(":5173", r))
}
