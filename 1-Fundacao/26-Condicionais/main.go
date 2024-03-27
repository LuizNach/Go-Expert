package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello\n")

	a := 10
	b := 20

	if a > b {
		println("a é maior que b")
	} else {
		println("b é maior que a")
	}

	if a == 1 {
		println("a é 1")
	} else if b == 20 {
		println("O golang tem else if sim!")
	}

	// condicoes booleanas: and &&, or ||
	//Para bitwise operations Reference: https://yourbasic.org/golang/bitwise-operator-cheat-sheet/
	a = 3
	b = 2
	c := 1
	if a > b && b > c {
		println("a>b>c")
	}

	// switch
	// Reference https://go.dev/tour/flowcontrol/9
	switch a {
	case 3:
		{
			fmt.Println("Entrou no 1.")
		}
	case 2:
		{
			fmt.Println("Entrou no 2.")
		}
	default:
		fmt.Println("Entrou no default")
	}

	// switch can use the walrus operator
	switch value := a > b; value {
	case true:
		fmt.Printf("Deu true no switch:%v\n", value)
	case false:
		fmt.Printf("Deu false no switch:%v\n", value)
	}
}
