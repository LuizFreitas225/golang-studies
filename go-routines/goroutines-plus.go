package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // Define o número máximo de CPUs que podem ser usadas pelo programa

	fmt.Println("Começou")

	go func() {
		for {
		}
	}()

	fmt.Println("Terminou")
}
