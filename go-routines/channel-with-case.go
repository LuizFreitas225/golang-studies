package main

import "time"

//Main é a primeira "trhead"

func main() {

	hello := make(chan string)
	go func() {
		time.Sleep(time.Second * 2) // Simulando um atraso de 2 segundos antes de enviar a mensagem
		hello <- "Hello, World!"
	}()

	//Nesse caso não haverá espera da goroutine, pois cairemos no caso default do select

	select {
	case x := <-hello:
		println(x)
	default:
		println("Nenhuma mensagem recebida")

	}
}
