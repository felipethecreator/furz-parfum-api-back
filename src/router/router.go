package router

import (
	"github.com/gorilla/mux"
)

// Gera um novo router com as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
