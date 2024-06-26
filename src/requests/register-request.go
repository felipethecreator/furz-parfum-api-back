package requests

import (
	"context"
	"encoding/json"
	"net/http"
	"ps-backend-Matheus-Musashi/src/database"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decode the JSON request body and store it in the user variable
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get a reference to the "users" collection
	collection := database.GetUserCollection()

	// Check if the username already exists
	var existingUser User
	err = collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Insert the user document into the "users" collection
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Set the HTTP status as 201 (created)
	w.WriteHeader(http.StatusCreated)

	// Encode the user variable to JSON and write it to the HTTP response
	json.NewEncoder(w).Encode(user)
}
