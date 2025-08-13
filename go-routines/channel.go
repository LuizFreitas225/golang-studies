package main

import "time"

//Main é a primeira "trhead"

func main() {
	// Criando um canal para comunicação entre goroutines
	// O canal é do tipo string, permitindo enviar e receber mensagens de texto.
	// Os canais são usados para sincronizar goroutines e passar dados entre elas.

	hello := make(chan string)

	//Essa trhead anônima  é a segunda "trhead"
	//	Iniciando uma goroutine anônima que envia uma mensagem para o canal hello
	// O <-  é usado para enviar dados para o canal.
	go func() {
		time.Sleep(time.Second * 2) // Simulando um atraso de 2 segundos antes de enviar a mensagem
		hello <- "Hello, World!"
	}()

	// Recebendo a mensagem do canal hello
	// O <- é usado para receber dados do canal.
	// Isso bloqueia a execução até que uma mensagem esteja disponível no canal.
	// Isso garante que a mensagem seja recebida antes de continuar a execução do programa.

	result := <-hello

	println(result)
}
