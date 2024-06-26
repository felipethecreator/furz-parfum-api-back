package components

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"ps-backend-Matheus-Musashi/src/database"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// função HTTP handler que lida com requisições de registro de usuários
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decodifica o corpo da requisição JSON e armazena os dados na variável user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Printf("Decoded user: %+v\n", user)
	// Obtém uma referência para a coleção "users" no banco de dados "userdb"
	collection := database.GetUserCollection()

	// Insere o documento user na coleção "users"
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Define o status HTTP como 201(created)
	w.WriteHeader(http.StatusCreated)

	// Codifica a variável user em JSON e escreve na resposta HTTP
	json.NewEncoder(w).Encode(user)
}
