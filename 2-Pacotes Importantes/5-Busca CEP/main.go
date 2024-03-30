package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func (c ViaCEP) prettyPrint() {
	fmt.Printf("cep: %v\n", c.Cep)
	fmt.Printf("logradouro: %v\n", c.Logradouro)
	fmt.Printf("complemento: %v\n", c.Complemento)
	fmt.Printf("bairro: %v\n", c.Bairro)
	fmt.Printf("localidade: %v\n", c.Localidade)
	fmt.Printf("uf: %v\n", c.Uf)
	fmt.Printf("ibge: %v\n", c.Ibge)
	fmt.Printf("gia: %v\n", c.Gia)
	fmt.Printf("ddd: %v\n", c.Ddd)
	fmt.Printf("siafi: %v\n", c.Siafi)
}

func main() {

	// os.Args é um slice de strings de todo o comando que iniciou a funcao comecando pelo nome do programa
	// fmt.Printf("Lista de argumentos:\n")
	// for index, argumento := range os.Args[:] {
	// 	fmt.Printf("Argumento %v é: %v\n", index, argumento)
	// }

	println("-----------------------------------")

	for _, cep := range os.Args[1:] { // não pegamos o proprio nome do programa
		fmt.Printf("CEP é: %v\n", cep)
		url := "https://viacep.com.br/ws/" + cep + "/json/"
		response, err := http.Get(url)

		if err != nil {
			fmt.Printf("Nao foi possivel buscar esse CEP: %v\n", cep)
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisicao: %v\n", err)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("Para o CEP: %v - Nao foi possivel acessar a url: %v\n", cep, url)
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}

		println("A resposta em json:")
		println(string(body))

		dado := ViaCEP{}
		err = json.Unmarshal(body, &dado)

		if err != nil {
			fmt.Printf("Para o CEP %v não foi possivel converter o dado apropriadamente\n", cep)
			fmt.Fprintf(os.Stderr, "Erro ao fazer a deserializacao: %v\n", err)
		}

		dado.prettyPrint()

		file, err := os.Create("cep-" + cep + ".txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo de saida cep+%v.txt: %v\n", cep, err)
		}

		_, err = file.WriteString(fmt.Sprintf("CEP: %v, UF: %v, Cidade: %v, Localidade: %v\n", cep, dado.Uf, dado.Localidade, dado.Logradouro))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer escrita no arquivo de saida cep+%v.txt: %v\n", cep, err)
		}

		println("-----------------------------------")
	}

}
