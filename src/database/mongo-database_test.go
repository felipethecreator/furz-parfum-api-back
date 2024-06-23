package database

import (
	"testing"
)

func TestConnectMongoDB(t *testing.T) {
	// Define a URI de conexão com o MongoDB Atlas
	uri := "mongodb+srv://felipe:1234@furz-parfum.jwzwsza.mongodb.net/?retryWrites=true&w=majority&appName=Furz-parfum"

	// Chama a função ConectarMongo passando a URI e captura o cliente MongoDB e qualquer erro retornado
	client, err := ConectarMongo(uri)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if client == nil {
		t.Fatal("Expected client to be non-nil")
	}
}
