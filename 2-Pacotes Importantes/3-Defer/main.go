package main

import (
	"fmt"
	"io"
	"net/http"
)

func funcao1() {
	fmt.Printf("Comecou a funcao1\n")
	defer fmt.Printf("Antes da funcao 2\n")
	funcao2()
	defer fmt.Printf("Depois da funcao 2\n")
	fmt.Printf("Terminou a funcao 1\n")

	/*
	   Comecou a funcao1
	   ...
	   Terminou a funcao 1
	   Depois da funcao 2
	   Antes da funcao 2
	*/
}
func funcao2() {
	fmt.Printf("Comecou a funcao 2\n")
	defer fmt.Printf("Antes da funcao 3\n")
	funcao3()
	defer fmt.Printf("Depois da funcao 3\n")
	fmt.Printf("Terminou a funcao 2\n")

	/*
	   Será exibido:
	   Comecou a funcao 2
	   ...
	   Terminou a funcao 2
	   Depois da funcao 3
	   Antes da funcao 3
	*/
}

func funcao3() {
	/*
	   Um bom detalhe é que o defer é empilhado e suas exeucoes são LIFO
	*/

	fmt.Printf("Comecou a funcao 3\n")
	defer fmt.Printf("Teste1\n")
	defer fmt.Printf("Teste2\n")
	defer fmt.Printf("Teste3\n")
	fmt.Printf("Terminou a funcao 3\n")

	/*
	   Será exibido:
	   Comecou a funcao 3
	   Terminou a funcao 3
	   Teste3
	   Teste2
	   Teste1
	*/
}

func requisicaoweb() {
	requisicao, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Erro ao realizar a requisição HTTP:", err)
		return
	}
	defer requisicao.Body.Close() // Aqui garantimos que após a função requisicaoweb encerrar será executado requisicao.Body.Close()

	response, err := io.ReadAll(requisicao.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return
	}

	fmt.Printf("O resultado da requisicao convertendo o stream de bytes para string:\n")
	fmt.Printf("%v\n", string(response))
}

func main() {
	/*
	   Defer é utilizado quando que quero atrasar mas garantir que um passo seja executado.

	   O defer em Go é uma declaração que adia a execução de uma função até que a função envolvente seja concluída.
	   Geralmente, é usado para garantir que determinadas operações sejam realizadas antes que a função retorne, como
	   fechar um arquivo ou liberar recursos. Isso ajuda a manter o código limpo e organizado, garantindo que as tarefas
	   necessárias sejam executadas no momento apropriado.

	   Os defers em uma função em Go são executados em ordem inversa, ou seja, o último defer a ser declarado será o
	   primeiro a ser executado (LIFO - Last In, First Out). Isso significa que o defer mais recente será executado antes
	   dos defers declarados anteriormente.
	*/

	println("-------------------------------------")

	funcao1()

	println("-------------------------------------")
}
