package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Hello\n")

	ctx := context.Background()                               // Criamos um contexto em background para poder ser utilizado
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond) //a partir desse contexto de background adicionamos uma regra de tomeout

	defer cancel() // o retorno da .WithTimeout retorna duas coisas a variavel de contexto e a funcao que cancela todos os recursos
	// que tiverem associados a esse contexto.

	// Podemos criar uma contrução de uma requisição http com um determinado um contexto. A requisição aqui
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)

	if err != nil {
		fmt.Printf("Falhou no ponto 1:\n")
		return
	}

	// Sem precisar criar um http client podemos usar o default client do go para executar a requisicao http
	resp, err := http.DefaultClient.Do(req) // Aqui ao executar uma requisicao com contexto, ela pode atingir a regra do contexto. No nosso caso é o

	if err != nil {
		fmt.Printf("Falhou no ponto 2:\n")
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Falhou no ponto 3:\n")
		return
	}

	fmt.Printf("Resposta do body:\n%s", body)
}
