package main

import (
	"fmt"
)

type Cliente struct {
	Nome  string
	Saldo int
}

func (c Cliente) andou() {
	c.Nome = "Luiz Nach" // estamos alterando o campo Nome da sctruct passada como copia para funcao andou
	fmt.Printf("Dentro da funcao, %s andou...\n", c.Nome)
}

func (c *Cliente) andou_com_referencia() {
	c.Nome = "Luiz Nach" // estamos alterando o campo Nome da sctruct passada como creferencia para funcao andou
	fmt.Printf("Dentro da funcao, %s andou com referencia...\n", c.Nome)
}

func (c Cliente) simulacao_bancaria(Valor_adicionar int) int {
	c.Saldo += Valor_adicionar
	fmt.Printf("Dentro da simulacao bancaria, o novo saldo é: %v\n", c.Saldo)
	return c.Saldo
}

func (c *Cliente) simulacao_bancaria_ref(Valor_adicionar int) int {
	c.Saldo += Valor_adicionar
	fmt.Printf("Dentro da simulacao bancaria Ref, o novo saldo é: %v\n", c.Saldo)
	return c.Saldo
}

// Uma boa pratica em go é a funcao geradora de novas instancias de struct
func NewClient() *Cliente {
	// Aqui a unica logica de negocio é que o cliente deve ser sempre criado com saldo 0
	endereco_de_memoria := &Cliente{Saldo: 0}
	fmt.Printf("Fazendo o NewClient com %T %v %v\n", endereco_de_memoria, endereco_de_memoria, *endereco_de_memoria)
	return endereco_de_memoria
}

func main() {
	fmt.Printf("Hello\n")

	usuario := Cliente{Nome: "Luiz"}

	// Chamamos a funcao andou, e queremos saber se apos a execucao da funcao os campos da sctruct foram alterados
	usuario.andou()
	fmt.Printf("Apos a execucao de andou\nO valor da struct é: %v\n", usuario)

	println("---------------------------")

	// Chamamos a funcao andou_com_referencia,
	// e queremos saber se apos a execucao da funcao os campos da sctruct foram alterados
	usuario.andou_com_referencia()
	fmt.Printf("Apos a execucao de andou com referencia\nO valor da struct é: %v\n", usuario)

	//Assim concluimos que utilizamos
	// funcoes com copia para quando queremos utilizar os valores dos campos mas nao altera-los no struct original
	// funcoes com referencia para quando queremos utilizar os valores dos campos mas queremos altera-los no struct original

	println("---------------------------")

	fmt.Printf("Antes da simulacao bancaria\n")
	fmt.Printf("O saldo do cliente é: %v\n", usuario.Saldo)
	usuario.simulacao_bancaria(10000)
	fmt.Printf("Depois da simulacao bancaria\n")
	fmt.Printf("O saldo do cliente é: %v\n", usuario.Saldo)

	println("---------------------------")

	fmt.Printf("Antes da simulacao bancaria por referencia\n")
	fmt.Printf("O saldo do cliente é: %v\n", usuario.Saldo)
	usuario.simulacao_bancaria_ref(10000)
	fmt.Printf("Depois da simulacao bancaria por referencia\n")
	fmt.Printf("O saldo do cliente é: %v\n", usuario.Saldo)

	println("---------------------------")

	fmt.Printf("Uma boa pratica em go, são as funcoes geradoras de novas structs\n")
	var usuario2 *Cliente = NewClient()
	var usuario3 *Cliente = NewClient()

	fmt.Printf("Usuario 2 criado no endereco: %v com valores: %v e apontando para conteudo: %v\n", &usuario2, usuario2, *usuario2)
	fmt.Printf("Usuario 3 criado no endereco: %v com valores: %v e apontando para conteudo: %v\n", &usuario3, usuario3, *usuario3)
	fmt.Printf("Voce pode reparar que o ponteiro esta em um endereco\nmas quando printamos %%v do valor no ponteiro recebemos uma interpretacao da estrutura como &{ 0}\n")
	fmt.Printf("Para imprimirmos corretamente o valor de endereco contido dentro de um ponteiro precisamos do %%p\n")
	// Reference: https://stackoverflow.com/questions/56112289/how-to-print-the-address-of-struct-variable-in-go
	fmt.Printf("Usuario 2 criado no endereco: %v com valores: %p e apontando para conteudo: %v\n", &usuario2, usuario2, *usuario2)
	fmt.Printf("Usuario 3 criado no endereco: %v com valores: %p e apontando para conteudo: %v\n", &usuario3, usuario3, *usuario3)

	// Para saabermos que o NewClient esta gerando enderecos de memoria diferentes para novas estruturas podemos
	if usuario2 == usuario3 {
		fmt.Printf("usuario2 == usuario3\n")
	} else {
		fmt.Printf("usuario2 != usuario3\n")
	}
}
