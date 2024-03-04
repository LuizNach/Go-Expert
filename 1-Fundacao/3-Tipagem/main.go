package main

// Podemos criar nossos proprios tipos com a palavra reservada type
type ID int

func main() {
	var (
		variavel ID = 1
	)
	println(variavel)
}
