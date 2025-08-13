package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// Inicia o servidor HTTP
	log.Println("Iniciando servidor na porta :8080")
	http.HandleFunc("/", home)
	// Inicia o servidor e captura possíveis erros
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("Iniciou minha request")
	defer log.Println("Requisição finalizada")
	// Simula um processamento que pode levar tempo
	//Se o cliente cancelar a requisição antes do x segundos, o contexto será cancelado
	//Caso contrário, a requisição será processada normalmente
	select {

	//A requisição do cliente tem o timeout de 3s
	//Como aqui definimos 5s para concluir a requisã́o com sucess ela será cancelada antes pelo cliente
	//Aumentando o timeout no client pra mais de 5s espera se o caso de sucesso da requisição

	case <-time.After(time.Second * 5):
		log.Println("Requisição processada com sucesso")
		w.Write([]byte("Requisição processada com sucesso"))
	case <-ctx.Done():
		log.Println("Requisição cancelada pelo cliente")
		http.Error(w, "Requisição cancelada", http.StatusRequestTimeout)
	}
}
