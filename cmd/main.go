package main

import (
	//pacote CORS -> permite requisições de diferentes origens (ajuda na implementação do frontend)
	"github.com/rs/cors"
	"log"
	"net/http"
	"ps-backend-felipe-rodrigues/src/components"
	"ps-backend-felipe-rodrigues/src/database"
)

func main() {
	database.IniciarMongo()

	// corsHandler é uma variável que guarda o middleware padrão do CORS
	// Este middleware adiciona cabeçalhos HTTP necessários para permitir requisições de diferentes origens
	corsHandler := cors.Default().Handler

	// Configura a rota para registro de usuários
	http.HandleFunc("/register", components.RegisterUser)

	log.Println("Servidor iniciado na porta 5173")

	// http.DefaultServeMux é o roteador HTTP padrão que encaminha requisições para os handlers registrados, como o RegisterUser
	go log.Fatal(http.ListenAndServe(":5173", corsHandler(http.DefaultServeMux)))

}

/*user := components.User{
	Name:        "Felipe Queiroz",
	DisplayName: "FelipePirocade2cm",
	Password:    "peidofrouxoquenemfazbarulhomasfede123",
}*/
