package main

// importamos o pacote nativo fmt para controlar fomratacao
// Esse pacote permite multiplas formas de apresentacao de dados
import (
	"fmt"
)

type ID int

func main() {
	var (
		a         = "alguma string"
		b         = 1
		c         = 0.1
		d float32 = 0.1
		e ID      = 1
	)

	fmt.Printf("O tipo das variaveis a é: %T o seu valor é: %v\n", a, a)
	fmt.Printf("O tipo das variaveis b é: %T o seu valor é: %v\n", b, b)
	fmt.Printf("O tipo das variaveis c é: %T o seu valor é: %v\n", c, c)
	fmt.Printf("O tipo das variaveis d é: %T o seu valor é: %v\n", d, d)
	fmt.Printf("O tipo das variaveis e é: %T o seu valor é: %v\n", e, e)
}
