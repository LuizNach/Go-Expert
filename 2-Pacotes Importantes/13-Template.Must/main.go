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

	/*
	   Os templates em Go são uma forma de gerar texto dinâmico, como HTML, a partir de um modelo pré-definido.
	   O pacote template é uma parte importante da biblioteca padrão de Go para lidar com templates.
	   O método Must é comumente utilizado para simplificar a manipulação de erros ao analisar templates.
	*/

	curso := Curso{"Go", 360}
	fmt.Printf("Curso: %s Carga Horária: %d\n", curso.Nome, curso.CargaHoraria)

	t := template.Must(template.New("Curso Template").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")) // template.Must

	/*
	   Sabemos eu o template.Parse pode retornar erro. Mas como não tratamos esse erro e já passamos para o template.Must?
	   O método template.Must é comumente utilizado em Go para simplificar o tratamento de erros ao analisar templates.
	   Ele recebe um template e um erro como argumentos e retorna o template, tratando automaticamente qualquer erro
	   retornado durante a análise do template.
	*/

	err := t.Execute(os.Stdout, curso) // Assim só apresentará o erro no momento de fazermos o execute

	if err != nil {
		fmt.Printf("Não foi possível fazer o template: %s\n", err)
	}

}
