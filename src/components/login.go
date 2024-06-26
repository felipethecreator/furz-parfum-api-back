package components

import (
	"context"
	"encoding/json"
	"net/http"
	"ps-backend-Matheus-Musashi/src/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// função HTTP handler que lida com requisições de login de usuários
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var foundUser User

	// Decodifica o corpo da requisição JSON e armazena os dados na variável user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Obtém uma referência para a coleção "users" no banco de dados "userdb"
	collection := database.GetUserCollection()

	// Busca o usuário no banco de dados pelo DisplayName e Password
	filter := bson.M{"name": user.Username, "password": user.Password}
	err = collection.FindOne(context.TODO(), filter).Decode(&foundUser)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Failed to login user", http.StatusInternalServerError)
		return
	}

	// Define o status HTTP como 200 (OK) e retorna o usuário encontrado como JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundUser)
}
