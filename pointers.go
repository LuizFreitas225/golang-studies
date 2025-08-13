package main

import "fmt"

func main() {

	//&x → retorna o endereço de memória da variável x
	//*p → acessa o valor que o ponteiro p apont

	x := 10
	p := &x // ponteiro para x
	z := p  // z também é um ponteiro para x

	fmt.Println("x =", x)
	fmt.Println("Endereço de x =", p)
	fmt.Println("Valor apontado por p =", *p)

	*z = 20 // altera o valor de x indiretamente
	// z = 20  Não funciona pois z é um ponteiro de int e não um int, ou seja, é *int e não int
	fmt.Println("Novo valor de x =", x)
}
