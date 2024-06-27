package main

import (
	"log"
	"net/http"
	"ps-backend-Matheus-Musashi/src/database"
	"ps-backend-Matheus-Musashi/src/requests"

	"github.com/rs/cors"
)

func main() {
	log.Println("Server is starting...")

	// Define a URI de conexão com o MongoDB Atlas
	uri := "mongodb+srv://tanaka:asdf@ftenergy.b9nzeat.mongodb.net/?retryWrites=true&w=majority&appName=FTenergy"

	// Chama a função ConectarMongo passando a URI e captura o cliente MongoDB e qualquer erro retornado
	client, err := database.ConectarMongo(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Define o cliente MongoDB globalmente
	database.Client = client

	// Middleware CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust this to restrict origins as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler

	// Configura a rota para registro de usuários usando a função RegisterUserHandler
	http.HandleFunc("/register", requests.RegisterUserHandler)
	// Configura a rota para login de usuários usando a função LoginUserHandler
	http.HandleFunc("/login", requests.LoginUserHandler)

	// Inicia o servidor HTTP na porta 8080
	log.Println("Server starting at 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(http.DefaultServeMux)))
}
