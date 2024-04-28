package main

import "fmt"

func main() {
	var number int = 10

	var ponteiroParaNumber *int = &number

	fmt.Println("Valor de number:", number)
	fmt.Println("Endereço de memória: ", &number)

	fmt.Println("valor apontado por ponteiro para number: ", *ponteiroParaNumber)
	fmt.Println("Endereço de memória armazenado por ponteiroParaNumber: ", ponteiroParaNumber)

	*ponteiroParaNumber = 50

	fmt.Println("Novo valor de number após a modificação", number)
}
