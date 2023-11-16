package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	// fmt.Println(config.StringConexaoBnaco)
	fmt.Println("Rodando a API")
	fmt.Println(config.SecretKey)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
