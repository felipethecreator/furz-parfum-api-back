package controllers

import (
	"net/http"
)

func RegisterUser(http.ResponseWriter, *http.Request) {
	w.Write([]byte("Criando o caba"))
}
func LoginUser(http.ResponseWriter, *http.Request) {
	w.Write([]byte("Logando o caba"))
} 