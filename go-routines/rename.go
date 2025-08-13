package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerID int, msg chan int, wg *sync.WaitGroup) {
	//defer é usado para garantir que a função Done seja chamada quando a goroutine terminar
	defer wg.Done()
	for res := range msg {
		fmt.Printf("Worker %d,  MSG: %d\n", workerID, res)
		time.Sleep(time.Second)
	}
}

func main() {
	// Criando um canal para comunicação entre goroutines
	msg := make(chan int)
	// Usando WaitGroup para aguardar a finalização de todas as goroutines
	// WaitGroup é usado para esperar que um conjunto de goroutines termine.
	// Ele permite que você adicione um contador de goroutines e aguarde até que todas tenham chamado Done.
	// Cada goroutine chama Done quando termina, e o main aguarda até que todas tenham terminado.
	wg := sync.WaitGroup{}
	wg.Add(10)

	go worker(1, msg, &wg)
	go worker(2, msg, &wg)
	go worker(3, msg, &wg)
	go worker(4, msg, &wg)
	go worker(5, msg, &wg)
	go worker(6, msg, &wg)
	go worker(7, msg, &wg)
	go worker(8, msg, &wg)
	go worker(9, msg, &wg)
	go worker(10, msg, &wg)

	for i := 0; i < 10; i++ {
		msg <- i
	}

	// Fechando o canal após enviar todas as mensagens
	// Fechar o canal é importante para sinalizar que não haverá mais mensagens enviadas.
	// Isso permite que as goroutines que estão lendo do canal saibam quando parar de ler
	close(msg)

	fmt.Println("Aguardando workers finalizarem")
	wg.Wait()
	fmt.Println("Finalizado")

}
