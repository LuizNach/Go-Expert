package main

import (
	"fmt"

	"github.com/LuizNach/Go-Expert/5-Packaging/6-Usando-workspaces/math"
)

func main() {
	m := math.Math{} // Dado que nao temos mais a e b exportados,
	//nao conseguimos incializar a struct Math iniacializando a e b
	fmt.Printf("Resultado: %d\n", m.Add()) // logo o resultado de Add sera o valor iniciado default de a e b

	m.MathInitilize(1, 2)
	fmt.Printf("Resultado apos MathInitilize: %d\n", m.Add())

	m2 := math.NewMath(3, 4) // aqui apesar de math2 ser uma struct de uso interno da biblioteca math, conseguimos exportar essa struct dela
	fmt.Printf("tipo de m2:%T Add:%d", m2, m2.Add())
}
