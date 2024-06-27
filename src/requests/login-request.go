package requests

import (
	"context"
	"encoding/json"
	"net/http"
	"ps-backend-Matheus-Musashi/src/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	var foundUser User

	// Decode the JSON request body and store it in the user variable
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get a reference to the "users" collection
	collection := database.GetUserCollection()

	// Search for the user in the database by username and password
	filter := bson.M{"username": user.Username, "password": user.Password}
	err = collection.FindOne(context.TODO(), filter).Decode(&foundUser)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Failed to login user", http.StatusInternalServerError)
		return
	}

	// Set the HTTP status as 200 (OK) and return the found user as JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundUser)
}
