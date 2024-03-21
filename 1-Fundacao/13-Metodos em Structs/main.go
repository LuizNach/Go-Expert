package main

import (
	"fmt"
)

type Cliente struct {
	Nome  string // Vai ser incializado com string vazia
	Idade int    // Vai ser incializado com 0
	Ativo bool   // Vai ser inicializado com false
}

// Assim fazemos com que liguemos essa funcao a struct Cliente.
// Assim para uma variavel usuario do tipo Cliente podemos invocar a funcao do jeito usuario.Imprimir()
func (c Cliente) Imprimir() {
	fmt.Printf("Os dados do usuario sao:\n")
	fmt.Printf("\tNome: %v\n", c.Nome)
	fmt.Printf("\tIdade: %v\n", c.Idade)
	fmt.Printf("\tAtivo: %v\n", c.Ativo)
}

// Aqui temos a mesma ligacao da struct com a funcao
// O detalhe é que implementamos como (c Cliente) e por isso toda vez que a funcao for invocada,
// passara uma copia de usuario para dentro da funcao para ser utilizada
func (c Cliente) Desativar() {
	c.Ativo = false
}

func (c Cliente) Ativar() {
	c.Ativo = true
}

// Aqui temos a mesma ligacao da struct com a funcao
// Porem aqui descrevemos de um jeito diferente durante a ligação com o caracter *
// Assim toda vez que a funcao for chamada, o que ela recebe é a referencia da estrutura passada e nao uma copia
// Logo dizemos que fazemos a modificacao in place, ou utilizamos a referencioa de memoria da struct passada
func (c *Cliente) AtivarInPlace() {
	c.Ativo = true
}

func main() {
	fmt.Printf("Hello\n")

	usuario := Cliente{
		Nome:  "Eu",
		Idade: 10,
		Ativo: false,
	}

	fmt.Printf("Fazendo a funcao imprimir com usuario.Imprimir()\n")
	usuario.Imprimir()

	println("--------------------------------")

	fmt.Printf("O usuario antes da ativacao: %v\n", usuario)
	fmt.Printf("Repare no campo usuario.Ativo: %v\n", usuario.Ativo)

	usuario.Ativar()

	fmt.Printf("O usuario apos a ativacao: %v\n", usuario)
	fmt.Printf("Repare no campo usuario.Ativo: %v\n", usuario.Ativo)

	println("--------------------------------")

	fmt.Printf("Huuumm...nao funcionou. \n")
	fmt.Printf("Deve ser pq apesar de termos attachado a funcao ativar a sctruct cliente, \n")
	fmt.Printf("ela modifica somente o valor passado por valor e nao por referencia\n")

	println("--------------------------------")

	fmt.Printf("Fazendo a ativação a funcao AtivarInPlace(com ponteiro)...\n")
	fmt.Printf("O usuario antes a ativacao in place: %v\n", usuario)
	fmt.Printf("Repare no campo usuario.Ativo: %v\n", usuario.Ativo)

	usuario.AtivarInPlace()

	fmt.Printf("O usuario apos a ativacao in place: %v\n", usuario)
	fmt.Printf("Repare no campo usuario.Ativo: %v\n", usuario.Ativo)
}
