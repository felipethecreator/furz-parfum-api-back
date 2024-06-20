package main

import (
	"fmt"
	"log"
	"net/http"
	"ps-backend-felipe-rodrigues/src/router"
)

//TODO: IMPLEMENTAR O MONGODB NESSE CÓDIGO

func main() {
	fmt.Println("Onion rings tamanho G (12 unidades")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5173", r))
}