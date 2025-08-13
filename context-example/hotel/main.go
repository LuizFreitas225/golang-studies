package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// Define um timeout de 3 segundos para a operação de reserva
	//Se a requisição demorar mais que 3 segundos, o contexto será cancelado
	//Isso é útil para evitar que a aplicação fique esperando indefinidamente por uma resposta
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	bookHotel(ctx)
}

// bookHotel tenta reservar um quarto de hotel, mas pode ser exceder se o contexto for cancealdo.
func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Tempo excedido para bookar o quarto")
	case <-time.After(time.Second * 5):
		fmt.Println("Quarto reservado com sucesso")
	}
}
