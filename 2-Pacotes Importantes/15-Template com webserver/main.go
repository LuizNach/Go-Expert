package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cursos := Cursos{
			{"Go", 360},
			{"Java", 350},
			{"NodeJs", 300},
		}

		t := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := t.Execute(w, cursos) // Aqui temos um detalhe que como aceitamos um writer como parametro entao podemos utilizar o proprio writer do httpresponse

		if err != nil {
			fmt.Printf("Não foi possível fazer o template: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	})

	http.ListenAndServe(":8080", nil)

}
