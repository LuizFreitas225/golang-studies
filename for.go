package main

func main() {
	//Lista de carros
	carros := []string{"Gol", "Uno", "Camaro", "Mustang"}

	//Formas de iterar sobre a lista com for

	println("Usando o tradicional com índice\n")
	for i := 0; i < len(carros); i++ {
		println(carros[i], i)
	}

	println("\nUsando o for range\n")
	//O for range devolve dois valores, o índice e o valor do elemento
	//Se não for usar o índice, pode usar o _ para ignorar o valor
	//Ex: for _, carro := range carros {
	// O range trabalaha com uma copia do slice, array ou map

	for i, carro := range carros {
		carro = carro + " - TESTE" // Isso não modifica a estrutura original
		println("Original Não Modificada: ", i, carros[i])
		println("Cópia  Modificada: ", i, carro)

	}

	println("\nUsando o for range com índice\n")
	//Se quiser modificar o valor , precisa usar o índice
	for i, carro := range carros {
		carros[i] = carros[i] + " - TESTE" // Isso modifica a estrutura original
		println("Original Modificado: ", i, carros[i])
		println("Cópia Não Modificada: ", i, carro)
	}

}
