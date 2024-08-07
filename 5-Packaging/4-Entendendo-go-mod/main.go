package main

import "github.com/google/uuid"

/**
* A mensagem:
* could not import github.com/google/uuid (no required module provides package "github.com/google/uuid")
* so sera removida quando conseguirmos baixar a dependencia externa.
* Para isso podemos rodar go get ou go mod tidy (somente depois que o codigo estiver escrito para saber que estamos utilizando a lib).

Se tentarmos rodar go run main.go tomaremos um erro:
main.go:3:8: no required module provides package github.com/google/uuid; to add it:
        go get github.com/google/uuid
Podemos seguir esse conselho ou podemos rodaro go mod tidy que vai baixar pacotes usamos,
eleminar pacotes nao usados e criar um go.mod que garante as versoes das dependencias com os hashes do ultimos commits
*/

func main() {
	println(uuid.New().String())
}
