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

	// Define a URI de conexão com o MongoDB Atlas
	uri := "mongodb+srv://felipe:1234@furz-parfum.jwzwsza.mongodb.net/mydatabase?retryWrites=true&w=majority&ssl=true"

	// Chama a função ConectarMongo passando a URI e captura o cliente MongoDB e qualquer erro retornado
	_, err := database.ConectarMongo(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5173", r))
}
