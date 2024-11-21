package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Reference: https://aprendagolang.com.br/carregando-configuracoes-de-arquivos-toml/
// https://toml.io/en/v1.0.0

type Config struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
	Db   string `toml:"db"`
}

var conf Config

func init() {

	// A Funcao init
	// é uma função reservado do Go que é chamada durante
	// a inicialização do programa, ou seja, antes da função main.
	data, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}

	if _, err := toml.Decode(string(data), &conf); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Vamos ler um arquivo de config toml")
	fmt.Println("Para isso vamos utiliza a lib BurntSushi")

	fmt.Printf("Host: %s\n", conf.Host)
	fmt.Printf("Port: %d\n", conf.Port)
	fmt.Printf("User: %s\n", conf.User)
	fmt.Printf("Pass: %s\n", conf.Pass)
	fmt.Printf("Db: %s\n", conf.Db)

}
