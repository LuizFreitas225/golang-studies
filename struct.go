package main

import (
	"encoding/json"
	"fmt"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   string
}

// Exibe é um método associado à struct Cliente.
// Métodos são funções que têm um receiver, que é uma instância da struct.
// Neste caso, o método Exibe imprime uma mensagem formatada com o nome do cliente.
// Isso demonstra como podemos adicionar comportamento às structs em Go.
// Métodos são uma forma de encapsular lógica relacionada a uma struct, permitindo que
// possamos operar sobre os dados contidos nela de maneira organizada e reutilizável.
// O uso de métodos em structs é uma prática comum em Go, pois permite que a lógica
// relacionada a uma entidade específica seja agrupada com os dados dessa entidade.
// O método Exibe pode ser chamado em qualquer instância da struct Cliente, e ele
// imprimirá uma mensagem personalizada com o nome do cliente.
func (c *Cliente) Exibe() {
	fmt.Printf("Exibindo Cliente pelo método: %s\n", c.Nome)
}

// Composição de structs
// ClienteInternacional é uma struct que contém um Cliente e um campo adicional para o país.
// Isso demonstra como podemos estender a funcionalidade de uma struct existente.
// A composição é uma forma de reutilização de código, onde uma struct pode conter outra struct
// como um de seus campos, permitindo que a nova struct herde os campos e métodos da struct original.
type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"` // COnfigurando o campo para ser serializado como "pais" em JSON
}

func main() {
	cliente := Cliente{
		Nome:  "João Silva",
		Email: "email@solinftec.com",
		CPF:   "123.456.789-00",
	}

	fmt.Println("Cliente:", cliente)

	cliente2 := Cliente{
		Nome:  "Maria Oliveira",
		Email: "mari@gmai.com",
		CPF:   "987.654.321-00",
	}

	fmt.Printf("Nome: %s, Email: %s, CPF: %s\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	clienteInternacional := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "John Doe",
			Email: "d@d.com",
			CPF:   "111.222.333-44",
		},
		Pais: "EUA",
	}

	//Chamando atributos e métodos de Cliente diretamente no ClienteInternacional
	// Isso demonstra como a composição permite acessar os campos da struct interna diretamente.

	fmt.Printf("Nome: %s, Email: %s, CPF: %s, País: %s\n", clienteInternacional.Nome,
		clienteInternacional.Email, clienteInternacional.CPF, clienteInternacional.Pais)
	// Chamando o método Exibe da struct ClienteInternacional, que herda o método Exibe de Cliente.
	clienteInternacional.Exibe()
	cliente.Exibe()

	// Serializando a struct Cliente para JSON
	clienteJson, err := json.Marshal(cliente)
	if err != nil {
		fmt.Println("Erro ao converter cliente para JSON:", err)
		return
	}
	fmt.Println("Cliente em JSON:", string(clienteJson))

	// Deserializando o JSON de volta para a struct Cliente
	jsonCliente4 := `{"Nome":"João Silva","Email":"email@solinftec.com","CPF":"123.456.789-00"}`
	var cliente4 Cliente
	json.Unmarshal([]byte(jsonCliente4), &cliente4)
	fmt.Println("Cliente convertido do JSON:", cliente4)

}
