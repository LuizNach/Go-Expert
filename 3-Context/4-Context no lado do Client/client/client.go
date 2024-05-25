package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Printf("Hello from client!\n")

	timeOutSeconds := 6

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOutSeconds)*time.Second)
	defer cancel()

	/*
	   Criamos um context com a regra de timeout para em expirar em 3 segundos.
	   Como o servidor só responde a requisição após passado 5 segundos,
	   a request desse client.go vai ser cancelada pelo seu context por causa da
	   sua regra de timeout em 3 segundos.

	   Caso alteremos o timeOutSeconds para 6 segundos, a requisição só vai expirar
	   se a requisiçpao demorar 6 segundos e como o servidor demora 5 segundos então
	   a requisição executará sem problemas.
	*/

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if err != nil {
		panic("Não conseguimos construir a requisição\n")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("Não foi possível executar a requisição\n")
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)

}
