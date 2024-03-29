package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://www.google.com")
	// response sera um pointer para a struct http.Response. Existe a struct da lib http que controle tudo e nela temos uma struct Response que possui tudo da resposta

	if err != nil {
		fmt.Println("Erro ao realizar a requisição HTTP:", err)
		return
	}
	fmt.Printf("A resposta foi:\n")

	fmt.Printf("%p\n", response) // podemos tratar como um ponteiro
	println("----------------------")
	fmt.Printf("%v\n", response) // isso vai retornar um &{} de uma estrutura
	println("----------------------")
	fmt.Printf("%v\n", *response) // aqui temos realmente o valor apontado

	println("----------------------")
	fmt.Printf("A response ainda é uma estrutura fechada e precisamos interpretar o conteúdo do corpo da resposta da requisicao\n")
	fmt.Printf("%v\n", response.Body)

	fmt.Printf("O pacote http no método GET necessita que usemos o io.ReadAll para interpretar a resposta do site requerido porque\n")
	fmt.Printf("a função http.Get retorna um objeto Response que contém o corpo da resposta como um io.ReadCloser.\n")
	fmt.Printf("Para ler e interpretar o conteúdo desse corpo de resposta, é necessário utilizar o io.ReadAll para ler todos os\n")
	fmt.Printf("dados do io.ReadCloser e convertê-los em bytes. Dessa forma, podemos processar e utilizar a resposta recebida do site requerido.\n")

	// Lendo e interpretando o corpo da resposta utilizando ioutil.ReadAll
	conteudo_do_site, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return
	}

	fmt.Printf("Todo o conteúdo do site agora em bytes\n")
	fmt.Printf("%v\n", conteudo_do_site)

	fmt.Printf("Todo o conteúdo do site agora convertido em string\n")
	// Convertendo os bytes do corpo da resposta em string e exibindo
	fmt.Printf("n%v\n", string(conteudo_do_site))

	// Nao podemos esquecer de fechar a conexão com o servidor
	response.Body.Close()
}
