package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello\n")

	var meu_dict_nao_init map[string]string

	fmt.Printf("Map/Dicionario nao inicializado: %v\n", meu_dict_nao_init)

	println("--------------------")

	var (
		meu_dict = map[string]int{"abc": 1, "def": 2}
	)

	println("[LOG] println de um map:", meu_dict)                             // vai ser o endereço de memória
	fmt.Printf("[LOG] fmt Printf como é o %%v do dictioanry: %v\n", meu_dict) // map[abc:1 def:2]

	println("--------------------")

	fmt.Printf("[LOG] Delete do elemeto def: %v\n", meu_dict["def"])
	delete(meu_dict, "def")
	fmt.Printf("[LOG] Após fazer o delete do elemento def: %v\n", meu_dict)
	fmt.Printf("[LOG] Importante lembrar que após deletar da pra acessar a chave de def: %v\n", meu_dict["def"])

	println("--------------------")

	meu_dict["ghi"] = 3
	fmt.Printf("[LOG] Adição do elemento ghi com valor %d\n", meu_dict["ghi"])
	fmt.Printf("[LOG] Após adição do ghi: %v\n", meu_dict)

	println("--------------------")

	println("[LOG] Fazendo um novo dict com a funcao make")
	novo_dict := make(map[string]int)                            // a funcao make faz o mapa inicial
	println("[LOG] println do novo_dict:", novo_dict)            // vai ser o endereço de memoria
	fmt.Printf("[LOG] fmt printf do novo dict: %v\n", novo_dict) // map[]

	println("--------------------")

	fmt.Printf("[LOG] Percorrendo as chaves do mapa: %v\n", meu_dict)

	for key, value := range meu_dict {
		fmt.Printf("[LOG] O valor da chave %v e o valor do valor:%v\n", key, value)
	}

}
