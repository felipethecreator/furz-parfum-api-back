package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IniciarMongo() {
	// URI de conexão com o MongoDB Atlas
	uri := "mongodb+srv://tanaka:asdf@ftenergy.b9nzeat.mongodb.net/?retryWrites=true&w=majority&appName=FTenergy"

	// Define as opções do cliente
	clientOptions := options.Client().ApplyURI(uri)

	// Cria um contexto com timeout de 10 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Tenta conectar ao MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Tenta fazer um ping no banco de dados para verificar a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB successfully!")
}
