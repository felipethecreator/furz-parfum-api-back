package controllers

import (
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando o caba"))
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logando o caba"))
} 