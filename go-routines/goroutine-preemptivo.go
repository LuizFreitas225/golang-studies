package main

import "time"

//Main é a primeira "trhead"

func main() {
	queue := make(chan int)
	go func() {
		i := 0
		for {
			queue <- i
			i++

			//Chegando aqui há uma espera até que a outra goroutine (main) leia o valor do canal
		}
	}()

	//Recebendo dados do canal
	//O for range lê valores do canal até que ele seja fechado.
	//Se o canal nunca for fechado, o loop continuará indefinidamente.
	// O uso do range permite processar cada valor enviado para o canal assim que ele estiver disponível.
	//Dessa forma as duas goroutines (main e a anônima) trabalaham de forma preemptiva
	//Ou seja, o Go gerencia a troca entre elas automaticamente, fazendo com que após o termino de cada for a outra goroutine seja executada

	for x := range queue {
		time.Sleep(time.Second)
		println("Recebido:", x)

		//Aqui há uma espera até que a goroutine anônima envie um novo valor para o canal
	}
}
