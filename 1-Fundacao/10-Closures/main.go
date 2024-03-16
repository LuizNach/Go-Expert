package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello\n")
	// Closures são funcoes anonimas que existem na capacidade do golang de que uma funcao pode conter outras funcoes

	fmt.Printf("teste 1: %v\n", teste1_soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Printf("teste 2: %v\n", teste2_soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func teste1_soma(lista ...int) int {
	total := 0
	for _, value := range lista {
		// println(index)
		total += value
	}
	return total
}

func teste2_soma(lista ...int) int {

	x := func() int {
		return 2 * teste1_soma(lista...)
	}()

	return x
}
