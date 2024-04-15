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

	fmt.Printf("Cep query param: %v\n", cepParam)

	// Buscamos no viacep.com
	cepStruct, err := BuscaCEP(cepParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Agora que temos nossa struct pronta podemos contruir a response de algumas maneiras

	// -- -- -- -- Jeito de fazer 1

	/*
		json.Marshal:
		A função json.Marshal é utilizada para serializar um objeto em JSON e retornar os dados serializados como um slice de bytes.
		É útil quando você deseja serializar um objeto em JSON e armazenar o resultado em uma variável ou passá-lo diretamente para o http.ResponseWriter.
		Geralmente é mais simples de usar e adequado para casos em que você precisa serializar um objeto em JSON de uma vez.
	*/

	// Dando certo, criamos o slice de bytes que sera a representacao do json
	jsonData, err := json.Marshal(cepStruct)
	// Em caso de erro alteramos o status para Internal Server Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Em caso de sucesso somente escrevemos o payload
	// Setamos header para status ok e application json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

	// -- -- -- -- Jeito de fazer 2

	/*
		Aqui nao fazemos a serialização da struct em slice de bytes, escrevemos diretamente seus bytes diretamente no io.Writer
		json.NewEncoder:
		O json.NewEncoder cria um encoder que permite escrever dados JSON diretamente em um io.Writer, como o http.ResponseWriter.
		É útil quando você deseja serializar dados em JSON e escrevê-los diretamente em um stream de saída, como um arquivo ou uma conexão de rede.
		É mais eficiente em termos de memória, pois permite a escrita incremental dos dados JSON.
	*/

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// // Aqui criamos um novo encoder que aponta para colocar sua saida como http.ResponseWriter
	encoder := json.NewEncoder(w)
	err = encoder.Encode(cepStruct) // E aqui que ele faz a conversão para bytes diretamente
	// Caso haja por algum motivo de nao ser possivel escrever no novo encoder retorna um Internal Server Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// -- -- -- -- Jeito de fazer 3

	// json.NewEncoder(w).Encode(jsonData) // Ousado, considera que não haverá erros para gerar um encoder e já escreve a struct na saida respondendo a requisição http

	/*
		        Nos exemplos acima, temos duas funções de handler diferentes que respondem a requisições HTTP com dados JSON.
		        A primeira função utiliza json.Marshal para serializar o objeto Person em JSON e escrever a resposta no http.ResponseWriter.
		        Já a segunda função utiliza json.NewEncoder para escrever os dados JSON diretamente no http.ResponseWriter.
			    Esses exemplos demonstram como implementar respostas HTTP completas utilizando as duas ferramentas de serialização JSON em Golang.
	*/
}

func BuscaCEP(cep string) (*ViaCEP, error) {

	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
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
