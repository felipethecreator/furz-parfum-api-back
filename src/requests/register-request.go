package requests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"ps-backend-felipe-rodrigues/src/components"
)

func RegistrarUsuario(user components.User) {
	// Converte o objeto User em JSON.
	body, _ := json.Marshal(user)

	// Cria uma requisição HTTP POST simulada com os dados do usuário.
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	// Define o cabeçalho Content-Type como application/json.
	req.Header.Set("Content-Type", "application/json")

	// Cria um ResponseRecorder para capturar a resposta da requisição.
	rr := httptest.NewRecorder()

	// Define o handler de registro de usuário.
	handler := http.HandlerFunc(components.RegisterUser)

	// Chama o handler de registro de usuário com a requisição simulada.
	handler.ServeHTTP(rr, req)

}
