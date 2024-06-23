package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// função que estabelece uma conexão com o MongoDB usando a URI fornecida
func ConectarMongo(uri string) (*mongo.Client, error) {
	// Define as opções do cliente MongoDB usando a URI fornecida
	clientOptions := options.Client().ApplyURI(uri)

	// Cria um contexto com timeout de 10 segundos para a operação de conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Garantir que o contexto seja cancelado ao final da função para liberar recursos
	defer cancel()

	// Tenta conectar ao MongoDB usando as opções do cliente e o contexto criado
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Tenta enviar um comando "ping" para o banco de dados para verificar a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		// Se o ping falhar, retorna nil e o erro
		return nil, err
	}

	// Se a conexão e o ping forem bem-sucedidos, armazena o cliente na variável global Client
	Client = client
	log.Println("Connected to MongoDB!")
	return Client, nil
}
