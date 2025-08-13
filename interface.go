package main

import "fmt"

// Interface Veiculo define um contrato que qualquer tipo de veículo deve cumprir.
// Ela exige que o tipo implemente os métodos buzina e getModelo.
// Isso permite que diferentes tipos de veículos possam ser tratados de forma polimórfica.
type Veiculo interface {
	buzina()
	getModelo() string
}

type Carro struct {
	Fabricante string
	Modelo     string
	Ano        int
}

type Moto struct {
	Fabricante string
	Modelo     string
	Ano        int
}

//Caso os métodos buzina e getModelo sejam implementados por Carro e Moto,
// eles podem ser usados de forma polimórfica através da interface Veiculo no caso contrario irá gerar erro de compilação.
// Isso significa que podemos tratar Carro e Moto como Veiculo, permitindo maior flexibilidade no código.
// Não é necessário o uso e 'implements' como no Java, o Go identifica automaticamente se todos os métodos da interface foram implementados.

func (c Carro) buzina() {
	fmt.Println("Carro buzinando: Beep Beep!")
}
func (m Moto) buzina() {
	fmt.Println("Moto buzinando: Brrr Brrr!")
}

func (c Carro) getModelo() string {
	return c.Modelo
}

func (m Moto) getModelo() string {
	return m.Modelo
}

type Pessoa struct {
	Nome    string
	idade   int
	Veiculo Veiculo
}

func (p Pessoa) andou() {
	fmt.Println(p.Nome, "andou de", p.Veiculo.getModelo())
	p.Veiculo.buzina()
}

func main() {

	carro := Carro{Fabricante: "Toyota", Modelo: "Corolla", Ano: 2020}

	//moto := Moto{Fabricante: "Honda", Modelo: "CB500", Ano: 2019}

	wesley := Pessoa{
		Nome:    "Wesley",
		idade:   30,
		Veiculo: carro,
	}

	wesley.andou()
}
