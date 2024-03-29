package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {

	/*
		Em Go, json.Marshal e json.Unmarshal são funções da biblioteca padrão que permitem serializar (transformar em JSON)
		e desserializar (transformar de JSON para um tipo de dados Go) respectivamente.

		Essas funções são muito úteis para trabalhar com dados em formato JSON em Go, permitindo a comunicação com APIs,
		sistemas externos e a persistência de dados em arquivos, por exemplo.
	*/
	conta := Conta{Numero: 1, Saldo: 0}

	// Como a gente pega uma struct e converte para um json
	/*
	   json.Marshal: Esta função é usada para transformar um valor Go em sua representação JSON.
	   Por exemplo, ao chamar json.Marshal em uma struct, ela será convertida em uma string JSON.
	*/
	resultado, err := json.Marshal(conta)
	if err != nil {
		fmt.Printf("Erro ao serializar a struct em json\n")
		return
	}

	// resultado será slice de bytes
	// err será error
	fmt.Printf("Resultado do Marshall de uma struct para json em bytes: %v\n", resultado)
	fmt.Printf("Resultado do Marshall de uma struct para json em string: %v\n", string(resultado))

	println("-------------------------------------")

	// Quando a gente quer receber a struct e ja converter para o valor em string do json, utilizamos um encoder
	// No golang o encoder recebe um valor, faz um encoding ja gravando em algum outro lugar: arquivo, stdout

	encoder := json.NewEncoder(os.Stdout) // aqui criamos um encoder de Json ja passando que estará configurado para sair direto no stdout

	fmt.Printf("Agora ocorre o Encode apontando para o Stdout. Se ele funcionar dever imprimir no terminal com um \\n no final da linha\n")
	err = encoder.Encode(conta) // toda vez que chamarmos o Encode saira o convertido em bytes para o terminal. Detalhe que ele ja converte para string

	// ou tb poderia utilizar json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		fmt.Printf("Erro ao fazer o encode")
		return
	}

	println("-------------------------------------")

	// trazendo um slice de bytes para popular uma estrutura zerada na memoria

	var jsonPuro []byte = []byte(`{"Numero":2,"Saldo":100}`)
	uma_conta_vazia := Conta{}

	// Unmarshal traduz o slice de bytes para um certo endereço na memoria de uma estrutura
	/*
	   json.Unmarshal: Esta função é usada para transformar uma string JSON em um valor Go.
	   Por exemplo, ao chamar json.Unmarshal em uma string JSON, ela será convertida de volta para um valor Go.
	*/
	err = json.Unmarshal(jsonPuro, &uma_conta_vazia)

	if err != nil {
		fmt.Printf("Erro ao fazer Unmarshal do slice de bytes: %v\n", string(jsonPuro))
		return
	}

	fmt.Printf("Dado o json: %v\n", string(jsonPuro))
	fmt.Printf("Apos o Unmarshal do slice de bytes para uma Conta: %v\n", uma_conta_vazia)

	println("-------------------------------------")

	// Se o slice de bytes não respeitar o contrato e tiver campos com nome errado
	var jsonErrado []byte = []byte(`{"n":1, "s":1000}`) // os campos n e s nao estao mapeado na struct Conta
	conta_para_json_errado := Conta{}

	err = json.Unmarshal(jsonErrado, &conta_para_json_errado)

	if err != nil {
		fmt.Printf("Erro ao fazer Unmarshal do slice de bytes errado: %v\n", string(jsonErrado))
	}

	fmt.Printf("Dado o json: %v\n", string(jsonErrado))
	fmt.Printf("Apos o Unmarshal do slice de bytes errado para uma Conta: %v\n", conta_para_json_errado) // Aqui vemos que o campo nao prrenchido continua com valor nulo

	println("-------------------------------------")

	// TAGS
	/*
		As tags em Go são metadados associados aos campos de uma estrutura (struct) que fornecem informações
		adicionais sobre como esses campos devem ser tratados. Elas são usadas principalmente para serialização e
		desserialização de dados, como no caso de trabalhar com JSON.

		É como se fossem os apelidos de cada campo de uma struct.

		Neste exemplo, as tags json:"nome" e json:"idade" indicam que, ao serializar ou desserializar uma instância
		da struct Pessoa para JSON, os campos Nome e Idade devem ser tratados com os nomes "nome" e "idade", respectivamente.

	*/

	type Pessoa struct {
		Nome  string `json:"nome"`
		Idade int    `json:"idade"`
	}

	cliente := Pessoa{}
	json_com_campos_tags := []byte(`{"nome":"Eu","idade":10}`)

	err = json.Unmarshal(json_com_campos_tags, &cliente)

	if err != nil {
		fmt.Printf("Erro ao fazer Unmarshal do slice de bytes errado: %v\n", string(json_com_campos_tags))
	}

	fmt.Printf("Dado o json: %v\n", string(json_com_campos_tags))
	fmt.Printf("O resultado de um json com campos que correspondem a tags é: %v\n", cliente)
	println("Um detalhe bem importante é que ao utulizar tags, a impressao padrao da struct muda.")

	err = encoder.Encode(cliente)

	if err != nil {
		fmt.Printf("Erro ao tentar imprimir na tela\n")
	}
	println("A exibição pelo decoder json dos campos Nome e Idade foram trocados por suas tags de json")

	println("-------------------------------------")
}
