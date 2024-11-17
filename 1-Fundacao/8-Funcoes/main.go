package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello\n")

	fmt.Printf("Retorno da funcao soma-simples: %v\n", soma(1, 2))

	println("-------------------------")

	value1, value2 := soma_e_maior_que_dez(5, 5)

	fmt.Printf("Retorno da funcao com multiplos retornos:: valor1: %v valor2: %v\n", value1, value2)

	println("-------------------------")

}

func uma_funcao_void() {
	// O detalhe dessa funcao e que ela nao determinou retorno, logo sera void () no value
	fmt.Printf("Executando uma funcao sem retorno")
}

func soma(a int, b int) int {
	// como a e b sao do mesmo tipo poderiamos declarar a funcao como func soma(a, b int) int
	return a + b
}

func soma_e_maior_que_dez(a, b int) (int, bool) {
	// golang consegue retorna mais de uma valor em seu retorno
	// na assinatura da funcao devemos colocar o par dos tipos do retorno entre ( )
	if a+b > 10 {
		return a + b, true
	} else {
		return a + b, false
	}

}
