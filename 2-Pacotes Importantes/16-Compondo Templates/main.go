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

	templates := []string{
		"template.html",
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("template.html").ParseFiles(templates...))
	// Os templates parsiados(header, content, footer) podem ser encontrados pelo template.html caso eu queira
	// No template.New passamos o nome do arquivo que queremos chamar
	// No ParseFIles passamos a lista de arquivos relacionadas a esse arquivo com ele incluso

	cursos := Cursos{
		{"Go", 360},
		{"Java", 350},
		{"NodeJs", 300},
	}
	err := t.Execute(os.Stdout, cursos)
	if err != nil {
		fmt.Printf("Não foi possível fazer o template: %s\n", err)
	}

}
