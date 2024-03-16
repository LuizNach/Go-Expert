package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello hello")
	//Funções variadicas são funções que levam um número variável de argumentos.
	fmt.Printf("O valor da soma total é: %v\n", soma(1, 2, 3, 4, 5, 6, 7, 8, 9))
	// Nao estamos passando uma estrutura de uma lista. Estamos passando varios argumentos
}

func soma(valores ...int) int {

	// O que determina que é uam função variádica é o termo reservado de ...
	total := 0
	for _, value := range valores {
		total += value
	}
	return total
}
