package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")

	funcUsesSomeValeuOnContext(ctx)
}

func funcUsesSomeValeuOnContext(ctx context.Context) {
	chave := ctx.Value("key")
	log.Printf("O valor da chave é: %s\n", chave)
}
