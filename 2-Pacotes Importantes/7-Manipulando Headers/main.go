package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", Root)

	http.HandleFunc("/buscaCep", buscaCep)

	http.ListenAndServe(":8080", nil)
}

func Root(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Requisicao, chegou em root\n")

	// Em r podemos pegar todos os parametros que vieram na requisicao
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!\n"))

}

func buscaCep(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Requisicao, chegou em busca Cep: %v\n", r.URL.Path)
	if r.URL.Path != "/buscaCep" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Cep param: %v\n", cepParam)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!\n"))
}
