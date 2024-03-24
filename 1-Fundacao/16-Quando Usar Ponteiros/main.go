package main

import (
	"fmt"
)

func soma(a, b *int) int {
	fmt.Printf("Primeiro ponteiro passado com valor: %v no endereco: %v\n", *a, a)
	fmt.Printf("Segundo ponteiro passado com valor: %v no endereco: %v\n", *b, b)
	return *a + *b
}

func soma2(a, b *int) int {
	*a = 1
	*b = 2

	fmt.Printf("Primeiro ponteiro passado com valor: %v no endereco: %v\n", *a, a)
	fmt.Printf("Segundo ponteiro passado com valor: %v no endereco: %v\n", *b, b)

	return *a + *b
}

func main() {
	fmt.Printf("Hello\n")

	fmt.Printf("Utilizamos ponteiros quando queremos possibilitar que outros contextos do código, outras funções, possam fazer alterações no dado.\n")
	fmt.Printf("Assim para fazermos alterações passamos o endereço da memoria exata\n")

	fmt.Printf("Fazendo testedo valor das variaveis antres da funcao que vai alterar elas\n")

	var x int = 2
	var y int = 3

	fmt.Printf("[Antes da execucao] Variavel x: %v\n", x)
	fmt.Printf("[Antes da execucao] Variavel y: %v\n", y)

	fmt.Printf("Fazendo a soma que altera os valores...\n")

	resultado := soma(&x, &y)

	fmt.Printf("Trazendo o resultado da funcao soma: %v\n", resultado)

	fmt.Printf("[Apos a execucao] Variavel x: %v\n", x)
	fmt.Printf("[Apos a execucao] Variavel y: %v\n", y)

	println("-------------------------")

	fmt.Printf("[Antes da execucao] Variavel x: %v\n", x)
	fmt.Printf("[Antes da execucao] Variavel y: %v\n", y)

	fmt.Printf("Fazendo a soma que altera os valores...\n")

	resultado = soma2(&x, &y)

	fmt.Printf("Trazendo o resultado da funcao soma: %v\n", resultado)

	fmt.Printf("[Apos a execucao] Variavel x: %v\n", x)
	fmt.Printf("[Apos a execucao] Variavel y: %v\n", y)
}
