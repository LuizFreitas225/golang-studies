package main

import (
	"fmt"
)

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

func main() {

	//Exemplo de if no Go
	num := 10
	if num > 0 {
		println("O número é positivo")
	} else if num < 0 {
		println("O número é negativo")
	} else {
		println("O número é zero")
	}

	//Exemplo de if com declaração
	//A variavel declarado no if só existe dentro do if e else
	if outroNum := -5; outroNum > 0 {
		println("O outro número é positivo")
	} else if outroNum < 0 {
		println("O outro número é negativo")
	} else {
		println("O outro número é zero")
	}

	//Exemplo de if com múltiplas condições
	idade := 18
	if idade >= 0 && idade < 18 {
		println("Menor de idade")
	}
	if idade >= 18 || idade < 0 {
		println("Maior de idade ou idade inválida")
	}

	//Exempo de if com declaração e tratamento de erro
	if res, err := dividir(10, 0); err != nil {
		println("Erro:", err.Error())
	} else {
		println("Resultado da divisão:", res)
	}

}
