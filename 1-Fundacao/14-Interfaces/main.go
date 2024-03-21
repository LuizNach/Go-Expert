package main

import (
	"fmt"
)

// Vamos dizer que eu quero falar que toda pessoa no meu sistema possui um metodo de consultar dados de seu perfil quando quiser.
// Uma maneira de falar isso é que temos que se algo é uma pessoa ela possui esse método de DadosCompletos()
// Em golang, fazemos isso definindo uma interface
type Pessoa interface {
	DadosCompletos()
}

// Agora para falarmos que podemos ter um a estrutura de cliente que se comporta como Pessoa, respeitando o contrato da interface Pessoa
// Devemos implementar a struct Cliente e attachar o metodo DadosCompletos nele
// Assim estaremos dizendo automaticamente que Cliente é uma pessoa, ou também que pode ser encarado como uma pessoa ao ser passado para funcoes
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// Attachando a estrutura Cliente a funcao DadosCompletos. A partir desse momento a estrutura Cliente pode ser encarada e passada como uma Pessoa
func (c Cliente) DadosCompletos() {

	fmt.Printf("Cliente\n\tNome: %v\n\tIdade: %v,\n\tAtivo: %v\n", c.Nome, c.Idade, c.Ativo)
}

// Vamos testar aqui criando uma funcao que recebe uma pessoa e vamos tentar passar um cliente
// Lembrar de voltar a esse conceito apos entender bem ponteiros pq interfaces com funcoes de ponteiros que se conectam a structs, funcionam de maneiras diferentes
func DadosCompletosUmaPessoaQueECliente(pessoa Pessoa) {
	fmt.Printf("Comecou a funcao\n")

	pessoa.DadosCompletos()

	fmt.Printf("Terminou a funcao\n")
}

func main() {
	fmt.Printf("Hello\n")

	user := Cliente{
		Nome:  "Nomezito",
		Idade: 1000,
		Ativo: false,
	}

	DadosCompletosUmaPessoaQueECliente(user)

}
