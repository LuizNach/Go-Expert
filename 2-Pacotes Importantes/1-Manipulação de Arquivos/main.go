package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("arquivo.txt")

	if err != nil {
		fmt.Printf("Deu ruim na criação do arquivo!\n")
		panic(err)
	}

	// tamanho, err := file.WriteString("Hello, mulecada\n") // escrevendo direto strings
	tamanho, err := file.Write([]byte("Hello galerinha bonita! Tudo belezinha?")) // escrevendo direto bytes

	if err == nil {
		fmt.Printf("O tamanho do arquivo gerado é: %d bytes.\n", tamanho)
	}

	file.Close()

	println("---------------------------------------")

	// Leitura
	/*
	   Podemos fazer diretamente a abertura do arquivo fazendo a leitura.
	*/
	input_file, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("O que vem do arquivo: %v e o que ele quer dizer: %s\n", input_file, input_file)
	fmt.Println(string(input_file)) // Como o que vem do arquivo é um slice de bytes devemos converter para string

	println("---------------------------------------")

	// Para o caso que o arquivo é muito grande, digamos que supera a nossa memória, temos que ler por partes.
	// Leitura em partes

	ponteiro_do_arquivo_aberto, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(ponteiro_do_arquivo_aberto) // utilizamos a lib de Buf Io para gerar um ponteiro de BufIoReader que possui a capacidade de ler aos poucos uma certa quantidade de bytes
	buffer := make([]byte, 10)                            //criamos um slice de bytes vazio de tamanho 10

	for { // loop infinito que só é encerrado com break
		n, err := reader.Read(buffer) // registra no slice de 10 em  10 bytes a cada iteracao retornando o n-esimo byte lido ate aquele momento
		// n sera 10 se conseguir ler até 10 bytes no reader, se houver menos bytes restantes teremos retornando o numero de quanto ele conseguiu fazer a leitura
		if err != nil { // Só apresenta erro quando chefa no End Of File
			fmt.Printf("O valor do error é: %v\n", err)
			break
		}
		fmt.Printf("%v tamanho em bytes: %v\n", string(buffer[:n]), n) // imprimir até o n-esimo byte no slice
	}

	// Remover arquivos
	err = os.Remove("arquivo.txt")
	if err != nil {
		fmt.Printf("Não foi possível remover o arquivo\n")
		panic(err)
	}

}
