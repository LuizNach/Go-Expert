package main

import (
	"fmt"
)

type Cliente struct {
	Nome  string // Vai ser incializado com string vazia
	Idade int    // Vai ser incializado com 0
	Ativo bool   // Vai ser inicializado com false
}

func main() {
	fmt.Println("Hello")

	usuario := Cliente{
		Nome:  "LuizNach",
		Idade: 1,
	}

	fmt.Printf("%T %v\n", usuario, usuario)
	fmt.Printf("Cliente Nome: %v, Idade: %v, Ativo: %v\n", usuario.Nome, usuario.Idade, usuario.Ativo)

	println("-------------------------------")

	usuario.Ativo = true
	fmt.Printf("Cliente após a mudança do ativo, Ativo: %v\n", usuario.Ativo)
	fmt.Printf("Cliente Nome: %v, Idade: %v, Ativo: %v\n", usuario.Nome, usuario.Idade, usuario.Ativo)
}
