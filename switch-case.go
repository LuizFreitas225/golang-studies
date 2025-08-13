package main

func main() {

	// Exemplo de switch case no Go
	numero := 2
	switch numero {
	case 1:
		println("Número é 1")
	case 2:
		println("Número é 2")
	case 3:
		println("Número é 3")
	default:
		println("Número não é 1, 2 ou 3")
	}

	// Exemplo de switch como if
	switch {
	case numero < 0:
		println("Número é negativo")
	case numero == 0:
		println("Número é zero")
	case numero > 0:
		println("Número é positivo")
	default:
		println("Número é inválido")
	}

}
