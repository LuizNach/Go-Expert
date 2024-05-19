package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Hello\n")

	c := http.Client{Timeout: time.Duration(2) * time.Second}

	resp, err := c.Get("http://google.com")

	if err != nil {
		panic("Deu erro na requisicao http")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic("Nao foi possivel resuperar o body da requisicao http")
	}

	fmt.Printf("Resposta foi:\n %s", body)

}
