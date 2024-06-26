package router

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"ps-backend-felipe-rodrigues/src/components"
	"ps-backend-felipe-rodrigues/src/database"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()

	// Define uma rota para o endpoint "/register" que usará a função RegisterUser do pacote components
	// Este endpoint aceitará apenas requisições do tipo POST
	r.HandleFunc("/register", components.RegisterUser).Methods("POST")

	// Define uma rota para o endpoint "/ping" que serve para verificar a conexão com o banco de dados
	// Este endpoint aceitará apenas requisições do tipo GET
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// Tenta enviar um comando de "ping" para o banco de dados MongoDB para verificar a conexão
		err := database.Client.Database("admin").RunCommand(r.Context(), bson.D{{"ping", 1}}).Err()
		if err != nil {
			http.Error(w, "Failed to ping database", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pinged your deployment. You successfully connected to MongoDB!"))
	}).Methods("GET")

	return r
}
