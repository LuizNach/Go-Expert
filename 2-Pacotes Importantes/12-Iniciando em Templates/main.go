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

func main() {
	curso := Curso{"Go", 360}
	fmt.Printf("Curso: %s Carga Horária: %d\n", curso.Nome, curso.CargaHoraria)
	tmp := template.New("Curso Template")
	tmp_errado, _ := tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraia}}") // Detalhe que sempre temos que acessar recurso com .
	err := tmp_errado.Execute(os.Stdout, curso)
	if err != nil {
		fmt.Printf("Não foi possível fazer o parse no template de curso: %s\n", err)
	}
	tmp_correto, _ := tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}") // Detalhe que sempre temos que acessar recurso com .
	err = tmp_correto.Execute(os.Stdout, curso)
	if err != nil {
		fmt.Printf("Não foi possível fazer o parse no template de curso: %s\n", err)
	}
}
