package main

import (
	"fmt"
)

func main() {

	var s []int

	fmt.Printf("O valor inicial do slice sem inicializacao: %v\n", s)

	println("-----------------------")
	println("Slice é essa estrutura que utiliza array por de baixo dos panos mas possui uma dinamica de reserva de memoria dinamica.")

	println("-----------------------")
	var new_slice = []int{1, 2, 3}
	fmt.Printf("New Slice: len=%d cap=%d valor=%v \n", len(new_slice), cap(new_slice), new_slice)

	println("-----------------------")
	new_slice = append(new_slice, 4)
	fmt.Println("Adicionando 4 ao new_slice")
	fmt.Printf("New Slice: len=%d cap=%d valor=%v \n", len(new_slice), cap(new_slice), new_slice)
	println("Eita po, a capacidade dobrou ao superar todos os espaços reservados do array")

	println("-----------------------")
	new_slice = append(new_slice, 5)
	fmt.Println("Adicionando 5 ao new_slice")
	fmt.Printf("New Slice: len=%d cap=%d valor=%v \n", len(new_slice), cap(new_slice), new_slice)

	println("-----------------------")
	new_slice = append(new_slice, 6)
	fmt.Println("Adicionando 6 ao new_slice")
	fmt.Printf("New Slice: len=%d cap=%d valor=%v \n", len(new_slice), cap(new_slice), new_slice)

	println("-----------------------")
	new_slice = append(new_slice, 7)
	fmt.Println("Adicionando 7 ao new_slice")
	fmt.Printf("New Slice: len=%d cap=%d valor=%v \n", len(new_slice), cap(new_slice), new_slice)

	println("Eita po, a capacidade dobrou ao superar todos os espaços reservados do array")

	// -- --
	println("-----------------------")
	println("Podemos tambem pegar somente uma parte do slice. Muito igual ao python!")

	fmt.Printf("New Slice[:0]: len=%d cap=%d valor=%v \n", len(new_slice[:0]), cap(new_slice[:0]), new_slice[:0])
	fmt.Printf("New Slice[:2]: len=%d cap=%d valor=%v \n", len(new_slice[:2]), cap(new_slice[:2]), new_slice[:2])
	fmt.Printf("New Slice[:5]: len=%d cap=%d valor=%v \n", len(new_slice[:5]), cap(new_slice[:5]), new_slice[:5])

	println("Podemos reparar que o len do vetor resultado muda porém o cap não se altera por ser uma capacidade reservada já e é propriedade do array original e não o resultante")

	fmt.Printf("New Slice[2:]: len=%d cap=%d valor=%v \n", len(new_slice[2:]), cap(new_slice[2:]), new_slice[2:])
	fmt.Printf("New Slice[6:]: len=%d cap=%d valor=%v \n", len(new_slice[6:]), cap(new_slice[6:]), new_slice[6:])

	println("Agora temos a capacidade diminuida por ignoramos o começo")
	println("-----------------------")
}
