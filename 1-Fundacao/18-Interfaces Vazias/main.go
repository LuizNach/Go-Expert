package main

import (
	"fmt"
)

// Implementar uma interface x com metodo y, vai fazer com que todas
// as structs que implementarem um metodo y, com mesmos parametros e assinatura
// estarao implementando a interface x
type x interface {
	y()
}

// Dado esse conhecimento, o que acontece quando temos uma interface w que nao possui
// metodos? Logo todas as structs do seu programa implementam essa interface
type w interface{}

// Isso é um trunfo para ser passado em qualquer funcao como um contrato porem deve ser utlizado com muita moderação

func showType(value interface{}) {
	// Essa função aceita qualquer parametro passado
	// Mas isso aumenta muito a chence de se fazer besteira no codigo
	fmt.Printf("Show Type of %v is %T\n", value, value)
}

func main() {
	fmt.Printf("A partir da versao 1.18 golang implementou Generics\n")

	var a = 10
	var b interface{} = "Hello World"

	fmt.Printf("O tipo da var a é: %T e o seu valor é: %v\n", a, a) // tipo int
	showType(a)                                                     // podemos passar uma variavel qualquer para um funcao que o contrato e uma interface vazia
	fmt.Printf("O tipo da var b é: %T e o seu valor é: %v\n", b, b) // tipo string
	showType(b)                                                     // podemos passar uma variavel que ja tenha sido declarada como um contrato de uma interface vazia
}
