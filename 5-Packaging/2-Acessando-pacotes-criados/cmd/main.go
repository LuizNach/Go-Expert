package main

import (
	"fmt"

	"github.com/LuizNach/Go-Expert/5-Packaging/2-Acessando-pacotes-criados/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Printf("Resultado: %d\n", m.Add())
}
