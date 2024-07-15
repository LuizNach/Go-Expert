package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {

	cotacao_ptr, err := getCotacao()
	if err != nil {
		log.Fatalf("Unable to get cotacao: %v", err)
	}

	const fileName string = "cotacao.txt"
	err = saveContentToFile([]byte(fmt.Sprintf("Dólar: %s\n", cotacao_ptr.Bid)), fileName)
	if err != nil {
		log.Fatalf("Unable to save content to file: %s\nError: %v", fileName, err)
	}
	log.Printf("Bid: %s\n", cotacao_ptr.Bid)

}

func getCotacao() (*CotacaoResponse, error) {

	const timeOutLimitMilisec = 300
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*timeOutLimitMilisec)
	defer cancel()

	const url = "http://localhost:8080/cotacao"
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error building http request:", err)
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Error executing http request: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Server didn't respond properly: %v", err)
		return nil, err
	}

	body_bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error parsing http request body:", err)
		return nil, err
	}

	var cotacao CotacaoResponse
	err = json.Unmarshal(body_bytes, &cotacao)
	if err != nil {
		log.Println("Error parsing json from body:", err)
		return nil, err
	}

	if cotacao.Bid == "" {
		return nil, errors.New("error on payload content: empty value")
	}
	return &cotacao, nil
}

func saveContentToFile(byteText []byte, fileName string) error {

	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(byteText)
	if err != nil {
		log.Println("Error writing byte slice to file:", err)
		return err
	}
	return nil
}
