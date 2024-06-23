package main

import (
	"log"
	"net/http"
	"ps-backend-felipe-rodrigues/src/database"
)

func main() {
	database.IniciarMongo()

	/*user := components.User{
		Name:        "Felipe Queiroz",
		DisplayName: "FelipePirocade2cm",
		Password:    "peidofrouxoquenemfazbarulhomasfede123",
	}*/

	log.Println("Servidor iniciado na porta 5173")

	go log.Fatal(http.ListenAndServe(":5173", nil))
}
