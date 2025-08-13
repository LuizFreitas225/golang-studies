package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	//res, _ := soma(10, 10) -> Forma de ignorar o erro

	res, err := soma(10, 10)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Resultado da soma:", res)
}

// Função com exemplo de retorno de erro
// Se a soma for maior que 10, retorna um erro
func soma(a, b int) (int, error) {
	res := a + b
	if res > 10 {
		return 0, errors.New("resultado maior que 10")

	}

	return res, nil
}
