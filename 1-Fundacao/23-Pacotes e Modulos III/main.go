package main

import (
	"fmt"

	"github.com/luiz.nach/pacotes-e-modulos-iii/matematica"
)

/**
* Exported names https://go.dev/tour/basics/3
* Exportar funcoes e variaveis
*
* Dentro de pacotes existe a possibilidade de variaveis e funcoes
* serem de uso exclusivo interno ou serem possiveis de serem utilizadas fora do pacote.
* Dentro de um pacote basta notar que:
*    funcoes podem ser exportadas se estiverem em capital (Primeira letra maiuscula)
*    variaveis é mesma regra ou seja letras minusculas e _ servem para uso interno
*    struct seguem a mesma regra, porem com o detalhe que acessar variaveis internas seguem a mesma regra
 */

func main() {
	fmt.Printf("Teste\n")

	println(matematica.Soma(1.5, 2.5))
	// println(matematica.soma(1.5, 2.5)) // undefined: matematica.soma // // Nao conseguimos importar uma funcao de uso interno do modulo matematica

	println("----------------------------------------")

	println("O valor da contante pi:", matematica.Pi)
	// println(_teste_underscore) // undefined // Nao conseguimos importar uma variavel de uso interno do modulo matematica

	println("----------------------------------------")

	println(matematica.Multiplicacao(1.5, 2.5))
	//println(matematica.multiplicacao(1.5, 2.5)) // undefined: matematica.multiplicacao //

	println("----------------------------------------")

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Printf("%v\n", carro) // detalhe que a string vazia de sub_marca será exibida
	carro.SetSubmarca("Uninho")

	// carro2 := matematica.Carro{Marca: "Fiat", marca: "teste"} // unknown field marca in struct literal of type matematica.CarrocompilerMissingLitField
	// println(carro2)
}
