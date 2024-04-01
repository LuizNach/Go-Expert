package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	http.HandleFunc("/buscaCep", buscaCep)

	http.ListenAndServe(":8080", nil)
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

func BuscaCEP(cep string) (*ViaCEP, error) {

	resp, err := http.Get("https://viacep.com.br/ws" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var cepStruct ViaCEP

	err = json.Unmarshal(body, &cepStruct)

	if err != nil {
		return nil, err
	}

	return &cepStruct, nil
}
