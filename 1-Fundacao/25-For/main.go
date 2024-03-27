package main

/*
Loopings

Para muitas linguagens temos jeitos muito diferentes para fazermos loops: foreach, for, do while, while, for in etc ou até pegada mais funcional com map, zip etc.
Dentro de Golang temos apenas o for
*/

import (
	"fmt"
)

func main() {

	fmt.Printf("Hello\n")

	// Um exemplo de um loop infinito
	// Para jobs, consumidores de filas/mensageria, rotinas constantes
	// for {
	//     //faça alguma coisa
	// }

	println("-----------------------------------")

	i := 0
	for i < 3 {
		println(i)
		i++
	}

	println("-----------------------------------")

	for i := 0; i < 3; i++ {
		println(i)
	}

	println("-----------------------------------")

	for i := range 3 {
		println(i)
	}

	println("-----------------------------------")

	numeros := []string{"abc", "def", "ghi"}

	for key, value := range numeros {
		fmt.Printf("Posicao: %v, Value: %v\n", key, value)
	}

	// Caso queiramos somente o value podemos utilizar o blank identifier no lugar de key
	for _, value := range numeros {
		fmt.Printf("Somente Value: %v\n", value)
	}

	println("-----------------------------------")

	um_mapa := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4} // lembrando que mapas em go nao manem a ordem das chaves, diferente do python
	um_mapa["f"] = 6
	um_mapa["e"] = 5
	um_mapa["g"] = 7

	for key, value := range um_mapa {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	println("-----------------------------------")

	var some_string string = "uma frase ai"

	for index, letter := range some_string {
		// println(index,letter) // so vai imprimir os inteiros do indice do asc de cada caracter
		fmt.Printf("Indice: %d,\tLetra em asc: %d,\tLetra: %c\n", index, letter, letter)
	}
}
