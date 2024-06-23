package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"ps-backend-felipe-rodrigues/src/components"
	"ps-backend-felipe-rodrigues/src/database"
)

func main() {
	log.Println("Server is starting...")

	// Define a URI de conexão com o MongoDB Atlas.
	uri := "mongodb+srv://felipe:1234@furz-parfum.jwzwsza.mongodb.net/mydatabase?retryWrites=true&w=majority&ssl=true"

	// Chama a função ConectarMongo passando a URI e captura o cliente MongoDB e qualquer erro retornado.
	client, err := database.ConectarMongo(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	database.Client = client

	user := components.User{
		Name:        "Felipe Rodrigues",
		DisplayName: "FelipeGOSTOSAO",
		Password:    "peidomolhadocheiodebosta",
	}

	// Converte o objeto User em JSON.
	body, _ := json.Marshal(user)

	// Cria uma requisição HTTP POST simulada com os dados do usuário.
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	// Define o cabeçalho Content-Type como application/json.
	req.Header.Set("Content-Type", "application/json")

	// Cria um ResponseRecorder para capturar a resposta da requisição.
	rr := httptest.NewRecorder()

	// Define o handler de registro de usuário.
	handler := http.HandlerFunc(components.RegisterUser)

	// Chama o handler de registro de usuário com a requisição simulada.
	handler.ServeHTTP(rr, req)

	log.Println("Servidor iniciado na porta 5173")

	go log.Fatal(http.ListenAndServe(":5173", nil))
}
