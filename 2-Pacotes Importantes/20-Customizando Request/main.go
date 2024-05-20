package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Printf("Hello, hello\n")

	c := http.Client{}

	req, err := http.NewRequest("GET", "http://google.com", nil)

	if err != nil {
		panic("Deu ruim na request")
	}

	req.Header.Set("Accept", "application/json")

	response, err := c.Do(req)

	if err != nil {
		panic("Deu ruim na response")
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic("Erro ao ler a resposta")
	}

	fmt.Printf("A resposta foi:\n%s", body)

}
