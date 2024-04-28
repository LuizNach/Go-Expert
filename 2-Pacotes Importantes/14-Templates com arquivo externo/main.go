package main

import (
	"fmt"
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	cursos := Cursos{
		{"Go", 360},
		{"Java", 350},
		{"NodeJs", 300},
	}

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, cursos)

	if err != nil {
		fmt.Printf("Não foi possível fazer o template: %s\n", err)
	}

}
