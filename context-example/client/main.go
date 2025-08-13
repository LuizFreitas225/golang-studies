package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	// Define um contexto com timeout de 3 segundos
	// Se a requisição demorar mais que 3 segundos, o contexto será cancelado
	//Canelando o contexto, a requisição HTTP será cancelada no server, evitando processe de forma desnecessária
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}

	// 	O que acontece se não chamar res.Body.Close()
	// O TCP connection não é liberado imediatamente.
	// O buffer que guarda os dados da resposta fica na memória.
	// Em aplicações com muitas requisições HTTP, isso acumula memória e conexões abertas,
	// levando a vazamentos ou até travamentos.
	defer res.Body.Close() // Sempre feche o corpo da resposta

	// Copia o corpo da resposta para a saída padrão
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Println("Erro ao ler corpo da resposta:", err)
	}
}
