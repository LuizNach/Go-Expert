package matematica

const Pi = 3.14

//const _teste_underscore = "Vamos ver no que que da..." // essa contante nao vai poder ser chamada fora da lib matematica

func Soma[T int | float32 | float64](a, b T) T {
	return soma(a, b)
}

func soma[T int | float32 | float64](a, b T) T { // essa funcao nao vai poder ser chamada fora da lib matematica
	return a + b
}

func multiplicacao[T int | float32 | float64](a, b T) T { // essa funcao nao vai poder ser chamada fora da lib matematica
	return a * b
}

func Multiplicacao[T int | float32 | float64](a, b T) T {
	return multiplicacao(a, b)
}

type Carro struct {
	Marca     string
	sub_marca string // essa propriedade nao vai poder ser acessada do lado de fora da lib matematica
}

func (c *Carro) SetSubmarca(nova_sub_marca string) {
	if c.Marca == "" {
		c.Marca = "sem Marca"
	}

	if c.sub_marca == "" {
		c.sub_marca = nova_sub_marca
	}

	println("O Carro de Marca:", c.Marca, "com a sub marca:", c.sub_marca)
}
