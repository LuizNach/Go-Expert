package main

/*
go get
- baixa uma dependencia especificada em seu argumento como caminho do github ou biblioteca padrao do go. ex: go get golang.org/x/exp/constraints
- após efetuar um go get, ele registra no go.mod a versão especifica
- após efetuar um go get, ele registra no go.sum as hash's de validaçpao do pacote com a versão especifica
*/

/*
go mod tidy
- Baixa as dependencias que vc estiver usando em seu projeto
- Retira do go.mod as depedenias que foram baixadas com go get mas não estão sendo utilizadas em seu projeto
- Ou seja mantém seu go.mod limpo
- Se você acicionou uma lib com go get, ela fica como //indirect no go.mod pq vc ainda não baixou ela localmente
para seu GOPATH então para instalar localmente necessitamos do go mod tody
*/

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Printf("Um novo uuid: %v", uuid.New())
}
