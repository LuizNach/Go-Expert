package main

import (
	"fmt"
)

// variavel const significa que uma vez atribuída ela não pode ser reatribuída
const a = "uma string a"

func main() {

	// a = "uma string b"
	// ./main.go:7:2: cannot assign to a (neither addressable nor a map index expression)
	println(a)

	// -- --

	// variaveis sao inicializadas por padrao com um valor inferido falsy do seu tipo.

	var (
		bool_sem_init   bool
		int_sem_init    int
		string_sem_init string
		float_sem_init  float64
	)

	println(bool_sem_init)   // No caso do bool   sera false
	println(int_sem_init)    // No caso do int    sera 0
	println(string_sem_init) // No caso do string sera ""
	println(float_sem_init)  // No caso do float  sera +0.000000e+000

	// --

	// Variaveis podem ser atribuidas colocando todas as palavras chave ou deixando que o compilador faça inferencia do tipo
	var minha_var_sem_inferencia float32 = 0.56
	minha_var_com_inferencia := 0.56

	// Para cada tipo ele possui uma inferencia padrão.
	fmt.Printf("%T", minha_var_sem_inferencia) // Tipo da variavel é float32
	fmt.Printf("%T", minha_var_com_inferencia) // Tipo da variavel é float64

	// Consigo atribuir e inferir o tipo para uma const?

	// const x := "aleluia" // nao rola
	// Reference: https://go.dev/tour/basics/15#:~:text=Constants%20can%20be%20character%2C%20string,declared%20using%20the%20%3A%3D%20syntax.

}

func nova_funcao() {
	// Escopo de variaveis
	// Reference: https://community.revelo.com.br/uso-de-variaveis-no-golang/
	// Em Golang, as variáveis podem ter escopos diferentes, dependendo de onde são declaradas.
	// Variáveis declaradas numa função, por exemplo, têm escopo local e só podem ser acessadas dentro daquela função.
	// Variáveis declaradas fora de qualquer função, por sua vez, têm escopo global e podem ser acessadas por qualquer função no pacote.

	println(a) // consegue recurar a por estar em escopo global
	//println(bool_sem_init) // nao consegue recuperar bool_sem_init por seu escopo ser local e definida em outra funcao
}
