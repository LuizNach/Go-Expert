package main

import (
	"fmt"

	"github.com/luiz.nach/curso-go/matematica"
)

// Ao criamos um go.mod para padronizar o repositorio como um possivel ponto de busca de lib para import
// Criamos go.mod com o comando go mod init <identificador-do-repositorio>
// Para registrarmos a possivel instalacao de dependencias externas rodamos go mod tidy, porem nao temos nenhuma dependencias externa no momento
// Assim podemos importar libs locais colocando o path a partir do identificador que criamos e o nome da lib a ser importada
// uma coisa interessante é que quando estmos utilizando um pacote de fora, ele baixa e guarda o pacote pelo seu identificador em GOPATH

// Agora com o nome do module escrito dentro do go.mod, podemos importar de dentro do projeto o package matematica e utilizar suas funcoes

func main() {
	fmt.Printf("O resultado de 2.5 + 3.5 é: %v\n", matematica.Soma(2.5, 3.5))
}
