package main

import (
	"fmt"
)

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string // Vai ser incializado com string vazia
	Idade    int    // Vai ser incializado com 0
	Ativo    bool   // Vai ser inicializado com false
	Endereco        // Aqui rola a composição
}

type Cliente2 struct {
	Address Endereco // Aqui não é composição. É uma variável/propriedade do tipo Endereco
}

func main() {
	fmt.Println("Hello")

	usuario := Cliente{
		Nome:  "LuizNach",
		Idade: 1,
	} // ususario incializado sem preencher a composição

	fmt.Printf("%T %v\n", usuario, usuario)
	fmt.Printf("Cliente Nome: %v, Idade: %v, Ativo: %v\n", usuario.Nome, usuario.Idade, usuario.Ativo)

	usuario.Logradouro = "nome da rua do usuario" // aqui podemos atribuir os atributos da composição como se fossem nativos
	usuario.Numero = 001
	usuario.Cidade = "cidade do usuario"
	usuario.Estado = "estado"

	println("-------------------------------")

	usuario.Ativo = true
	fmt.Printf("%T %v\n", usuario, usuario)
	fmt.Printf("Cliente\n\tNome: %v\n\tIdade: %v\n\tAtivo: %v\n", usuario.Nome, usuario.Idade, usuario.Ativo)
	fmt.Printf("\tLogradouro: %v\n\tNumero: %v\n\tCidade: %v\n\tEstado: %v\n", usuario.Logradouro, usuario.Numero, usuario.Cidade, usuario.Estado)

	println("-------------------------------")

	println("Incialização do usuario 2")
	usuario2 := Cliente{ // Aqui podemos inicializar passando a incialização da composição
		Nome:  "Isabella",
		Idade: 10,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "nome da rua",
			Numero:     100,
			Cidade:     "uma cidade",
			//Estado: "", // vamos deixar faltando estado para ver o que acontece
		},
	}

	fmt.Printf("%T %v\n", usuario2, usuario2)
	fmt.Printf("Cliente 2\n\tNome: %v\n\tIdade: %v\n\tAtivo: %v\n", usuario2.Nome, usuario2.Idade, usuario2.Ativo)
	fmt.Printf("\tLogradouro: %v\n\tNumero: %v\n\tCidade: %v\n\tEstado: %v\n", usuario2.Logradouro, usuario2.Numero, usuario2.Cidade, usuario2.Estado)

	println("-------------------------------")

	usuario3 := Cliente{}

	usuario3.Nome = "Cris"
	usuario3.Idade = 10
	usuario3.Ativo = true
	usuario3.Endereco.Cidade = "alguma cidade ai" // Aqui podemos acessar os atributos através da composição, talvez para deixar claro que aquela propriedade vem do resultado de uma composição
	usuario3.Endereco.Logradouro = "alguma rua ai"
	usuario3.Endereco.Numero = 100
	usuario3.Endereco.Estado = "algum estado ai"

	fmt.Printf("%T %v\n", usuario3, usuario3)
	fmt.Printf("Cliente 2\n\tNome: %v\n\tIdade: %v\n\tAtivo: %v\n", usuario3.Nome, usuario3.Idade, usuario3.Ativo)
	fmt.Printf("\tLogradouro: %v\n\tNumero: %v\n\tCidade: %v\n\tEstado: %v\n", usuario3.Endereco.Logradouro, usuario3.Endereco.Numero, usuario3.Endereco.Cidade, usuario3.Endereco.Estado)

	println("-------------------------------")

}
