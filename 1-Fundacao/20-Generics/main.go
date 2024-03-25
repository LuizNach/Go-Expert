package main

import (
	"fmt"
)

// ------------------------------------------------------------------------------------------------
// Generics vem para solucionar o problema de reescrever mesmas funcoes para tipos diferentes
func SomaInt(mapa map[string]int) int {
	var soma int
	for _, value := range mapa {
		soma += value
	}
	return soma
}

func SomaFloat(mapa map[string]float64) float64 {
	var soma float64
	for _, value := range mapa {
		soma += value
	}
	return soma
}

// Ao inves de ter dois codigos SomaInt e SomaFloat podemos ter um só
func SomaGenerics[T int | float64](mapa map[string]T) T {
	var soma T
	for _, value := range mapa {
		soma += value
	}
	return soma
}

// ------------------------------------------------------------------------------------------------
// Outra tecnica e usar uma constraint (condicao de restricao)
type Number interface {
	int | float64
}

// mas para utizarmos colocamos a interface constraint como generics
func SomaContraint[T Number](mapa map[string]T) T {
	var soma T
	for _, value := range mapa {
		soma += value
	}
	return soma
}

// ------------------------------------------------------------------------------------------------
// Agora se criarmos uma interface do tipo int ele pode nao ser reconhecido pelo tipo Number criado
type MyNumber int

// O tipo MyNumber2 se for usado como constraint, nao aceitara MyNumber apesar de ambos serem no fundo int
type MyNumber2 interface {
	int | float64
}

// Para resolver isso o go tem o caracter ~ para indicar que vai fazer a interpretacao do tipo na interface validar
// Assim variaveis MyNumber serao aceitas em constraints MyNumber3
type MyNumber3 interface {
	~int | ~float64
}

func SomaContraint2[T MyNumber2](mapa map[string]T) T {
	var soma T
	for _, value := range mapa {
		soma += value
	}
	return soma
}

func SomaContraint3[T MyNumber3](mapa map[string]T) T {
	var soma T
	for _, value := range mapa {
		soma += value
	}
	return soma
}

// ------------------------------------------------------------------------------------------------
// Agora um exercicio de comparacao.
// Aqui temos o exemplo da interface comparable, uma constraint nativa
// Existem outras constraints para ser usada em generics em libs nativas do golang

func Compara[T comparable](value1, value2 T) bool {
	// a interface comparable implementa varias tipos porem permite somente com igualdade ou diferente
	// operacoes como >, >=, <, <= não são permitidas
	// if value1 > value2 { //invalid operation: value1 > value2 (type parameter T is not comparable with >)compilerUndefinedOp
	//     return true
	// } else {
	//     return false
	// }

	if value1 != value2 {
		return true
	}
	return false

}

func main() {
	fmt.Printf("Hello\n")
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]float64{"a": 1.5, "b": 2.5, "c": 3.5}
	m3 := map[string]float64{"a": 1, "b": 1.5}
	println(SomaInt(m1))
	println(SomaFloat(m2))
	println(SomaGenerics(m3))
	println(SomaContraint(m1))
	println(SomaContraint(m2))
	println(SomaContraint(m3))

	println("-----------------------------------------------")

	m4 := map[string]MyNumber{"a": 1, "b": 2, "c": 3}
	fmt.Printf("A variavel mapa4 é do tipo: %T e possui valor: %v\n", m4, m4)

	// println(SomaContraint(m4)) // MyNumber does not satisfy Number (possibly missing ~ for int in Number)compilerInvalidTypeArg
	// println(SomaContraint2(m4)) // MyNumber does not satisfy MyNumber2 (possibly missing ~ for int in MyNumber2)compilerInvalidTypeArg
	println("SomaContraint3 consegue interpretar que o tipo MyNumber é no fundo int e um sub tipo de MyNumber3")
	println(SomaContraint3(m4)) // Aceito sem problemas graças ao ~

	println("-----------------------------------------------")
}
