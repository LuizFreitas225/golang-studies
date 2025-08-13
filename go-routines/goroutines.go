package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)

	}
}

func main() {
	// Iniciando duas goroutines que executam a função contador
	go contador("A")
	go contador("B")

	// Aguardando 10 segundos para garantir que as goroutines tenham tempo de executar
	time.Sleep(time.Second * 10)

	// contador("Sem go routine")
	// go contador("Com go routine")

	// fmt.Println("Hello 1")
	// fmt.Println("Hello 2")

	// time.Sleep(time.Second) // Espera 2 segundos para garantir que a goroutine termine antes de finalizar o programa
	// fmt.Println("Fim do programa")
}
