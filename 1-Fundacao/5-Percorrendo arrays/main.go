package main

import (
	"fmt"
)

func main() {
	var (
		minha_var any    // uma variaval sem tipo, por padrão incializada com nil
		meu_array [3]int // uma variaval vetor de int, com tamanho fixo não mutável
	)

	fmt.Printf("Tipo de minha_var é %T e seu valor é %v\n", minha_var, minha_var)
	fmt.Printf("Tipo de meu_array é %T e seu valor é %v\n", meu_array, meu_array)

	minha_var = 1
	meu_array[0] = 10
	meu_array[1] = 20
	meu_array[2] = 30
	fmt.Printf("Tipo de minha_var é %T e seu valor é %v\n", minha_var, minha_var)
	fmt.Printf("Tipo de meu_array é %T e seu valor é %v\n", meu_array, meu_array)

	for indice := 0; indice < len(meu_array); indice++ {
		meu_array[indice] = meu_array[indice] + 1
	}

	for i, v := range meu_array {
		fmt.Printf("Novo valor do indice %d no meu_array é %d\n", i, v)
	}

}
