package math

type Math struct {
	a int // Passamos a nao exporta a variavel a quando comecamos seu nome com letra minuscula
	b int // Passamos a nao exporta a variavel b quando comecamos seu nome com letra minuscula
}

func (m Math) Add() int {
	return m.a + m.b // Estando dentro do mesmo pacote, conseguimos acessar as variaveis a e b
}

func (m *Math) MathInitilize(x int, y int) {
	m.a = x
	m.b = y
}

type math2 struct { // struct nao exportavel
	a int
	b int
}

func NewMath(x int, y int) math2 { // com um metodo assim conseguimos retornar uma sctruct nao exportavel
	return math2{a: x, b: y} // instancia uma nova sctruct interna math2 para retornar
}

func (m math2) Add() int {
	return m.a + m.b
}
