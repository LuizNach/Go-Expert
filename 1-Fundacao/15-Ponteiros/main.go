package main

import (
	"fmt"
)

func TrocaPassagemPorValor(a, b int) {
	var temp = b

	fmt.Printf("Antes temp esta no endereco:%v valor:%v\n", &temp, temp)
	fmt.Printf("Antes a esta no endereco:%v valor:%v\n", &a, a)
	fmt.Printf("Antes b esta no endereco:%v valor:%v\n", &b, b)
	println("-------------------------------------")

	b = a
	a = temp

	fmt.Printf("Depois temp esta no endereco:%v valor:%v\n", &temp, temp)
	fmt.Printf("Depois a esta no endereco:%v valor:%v\n", &a, a)
	fmt.Printf("Depois b esta no endereco:%v valor:%v\n", &b, b)
	println("-------------------------------------")
}

func TrocaPassagemPorReferencia(a, b *int) {

	var temp int = *a
	// temp e um inteiro para guardar algum conteudo para ser efetuada a troca
	// como a é um ponteiro para int, fazendo *a estamos buscando o valor no endereco apontado por a
	// e armazenando uma copia em temp

	fmt.Printf("Antes temp esta no endereco:%v valor:%v\n", &temp, temp)
	fmt.Printf("Antes a aponta para endereco:%v valor:%v\n", a, *a)
	fmt.Printf("Antes b aponta para endereco:%v valor:%v\n", b, *b)
	println("-------------------------------------")

	*a = *b   // Quero que o [conteudo][do endereco de a] receba o conteudo do endereco de b
	*b = temp // Quero que o [conteudo][do endereco de b], receba o valor armazenado em temp

	fmt.Printf("Depois temp endereco:%v valor:%v\n", &temp, temp)
	fmt.Printf("Depois a aponta para endereco:%v valor:%v\n", a, *a)
	fmt.Printf("Depois b aponta para endereco:%v valor:%v\n", b, *b)
	println("-------------------------------------")
}

func main() {
	fmt.Printf("Hello\n")

	println("Todas as variaveis possuem valores e seus respetivos enderecos de memoria")
	println("Por exemplo inteiros")
	println("numero := 10")
	numero := 10
	fmt.Printf("A varivel numero tem dentro de si o valor: %v\n", numero)
	fmt.Printf("esta alocada em um certo local na memoria de endereco:%v\n", &numero)

	println()
	println("-------------------------------------")
	println()

	fmt.Printf("Para manipularmos enderecos de memoria utilizamos  um tipo de variavel especial chamada ponteiro.\n")
	fmt.Printf("Tal qual o inteiro tem o espaco de memoria de 32bits para armazenar numeros e portanto so pode receber numeros,\n")
	fmt.Printf("os ponteiros sempre tem o espaco para armazenar o hexadecimal do endereco de memoria e\n")
	fmt.Printf("sempre sao definidos com o tipo de variavel que vai ser armazenada naquele local de memoria\n")

	println()

	var pointer_de_int *int
	fmt.Printf("var pointer_de_int *int\n")
	fmt.Printf("A variavel pointer_de_int é um ponteiro para inteiros\n")
	fmt.Printf("Dentro dela sempre havera o valor de um endereco mas quando nao é incialziada sera nil: %v\n", pointer_de_int)

	println()

	fmt.Printf("Se colocarmos o endereco da variavel numero dentro de pointer_de_int podemos alterar numero por meio de pointertde_int\n")
	fmt.Printf("Essa e agrande sacada dos ponteiros, pq se vc tem acesso direto ao endereco vc pode alterar aquele local especifico em situacoes complexas\n")

	pointer_de_int = &numero
	fmt.Printf("pointer_de_int = &numero\n")
	fmt.Printf("A variavel pointer_de_int esta no endereco: %v\n", &pointer_de_int)
	fmt.Printf("A variavel pointer_de_int possui dentro si o valor: %v que nao por acaso e o endereco da var numero\n", pointer_de_int)
	fmt.Printf("Se pergarmos o conteudo da variavel pointer_de_int e interpretarmos buscando o que tem naquele endereco de memoria, encontramos: %v\n", *pointer_de_int)
	fmt.Printf("Que e o valor da varivael numero: %v\n", numero)

	println()

	println("-------------------------------------")
	println("-------------------------------------")
	println("-------------------------------------")

	println()

	println("Agora vamos tentar fazer os exercicios de trocar dois valores para entender um funcionamento de ponteiros")

	x := 10
	y := 20
	println("x := 10")
	println("y := 20")

	println("-------------------------------------")

	fmt.Printf("Antes de tudo:\n")
	fmt.Printf("x endereco:%v valor:%v\n", &x, x)
	fmt.Printf("y endereco:%v valor:%v\n", &y, y)

	println("-------------------------------------")

	fmt.Printf("Executando TrocaPassagemPorValor...\n")
	TrocaPassagemPorValor(x, y)

	println("-------------------------------------")

	fmt.Printf("Apos troca por valor => x: %v y:%v\n", x, y)

	println("-------------------------------------")

	fmt.Printf("Executando TrocaPassagemPorReferencia...\n")
	TrocaPassagemPorReferencia(&x, &y)

	println("-------------------------------------")

	fmt.Printf("Apos troca por referencia => x: %v y: %v\n", x, y)

	println("-------------------------------------")

}
