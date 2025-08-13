package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Fazendo uma requisição HTTP para um site inexistente
	// Em caso de erro, a função Get retornará um erro dentro do segundo valor, nessse caso o err
	res, err := http.Get("http://googddsfle.com.br")

	// Verificando se ocorreu um erro na requisição
	// Se err for diferente de nil, significa que houve um erro
	// Executar ação de tratamento de erro
	if err != nil {
		log.Fatal("Erro ao fazer a requisição")
		log.Fatal(err.Error())
		panic(err)
	}

	fmt.Println(res.Header)
}
