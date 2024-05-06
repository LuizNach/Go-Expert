package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

/*
Os pacotes html/template e text/template são pacotes de template do Go que oferecem funcionalidades para
geração de texto formatado a partir de templates. A diferença principal entre eles está na forma como
lidam com a segurança do conteúdo gerado.

O pacote html/template é utilizado para gerar HTML seguro, protegendo contra ataques de script entre
outras vulnerabilidades. Ele realiza a escape automática de variáveis inseridas nos templates, garantindo
que o conteúdo gerado seja seguro para ser renderizado em um navegador.

Já o pacote text/template é mais genérico e não realiza a escape automática de variáveis.
Ele é mais adequado para gerar texto simples, como por exemplo, arquivos de configuração, e não é
recomendado para gerar HTML diretamente, pois pode resultar em vulnerabilidades de segurança.

Portanto, a principal diferença entre os dois pacotes está na segurança do conteúdo gerado:
o html/template é mais seguro para gerar HTML, enquanto o text/template é mais genérico e não realiza
escape automática de variáveis.
*/

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(text string) string {
	return strings.ToUpper(text)
}

func batatinha(text string) string {
	// O seguinte texto vai ser impresso no meio do stdout
	// Caso o template esteja colocando a saida dos html no stdout, vai concorrer com o batatinhas
	fmt.Printf("Esse texto: %s passou por batatinha", text)
	return text
}

func main() {

	// Criamos o array de strings para guardar todos os arquivos html que se relacionam
	templates := []string{
		"template.html",
		"header.html",
		"content.html",
		"footer.html",
	}

	// Criamos um template com o nome do content.html
	t := template.New("content.html")

	// Aqui mapeamos funcoes que podem ser invocadas dentro dos templates para utilizarmos no meio da renderização do template
	t.Funcs(template.FuncMap{"ToUpper": ToUpper, "cenorinha": batatinha})

	// Aqui passamos toda a lista de arquivos html que utilizarao as funcoes passadas
	t = template.Must(t.ParseFiles(templates...))

	cursos := Cursos{
		{"Go", 360},
		{"Java", 299},
		{"NodeJs", 300},
	}

	// Aqui passamos o template para executar na variavel cursos, ou seja, vai pode iterar para renderizar multiplos elementos
	err := t.Execute(os.Stdout, cursos)
	if err != nil {
		fmt.Printf("Não foi possível fazer o template: %s\n", err)
	}

}
