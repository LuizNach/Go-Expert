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
var newGlobalConfig Config2

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

	data, err = os.ReadFile("config2.toml")
	if err != nil {
		panic(err)
	}

	_, err = toml.Decode(string(data), &newGlobalConfig)
	if err != nil {
		panic(err)
	}
}

func main() {

	// Para rodar a main.go use: go run .
	fmt.Println("Vamos ler um arquivo de config toml")
	fmt.Println("Para isso vamos utiliza a lib BurntSushi")

	fmt.Println("====================================")

	fmt.Printf("Host: %s\n", conf.Host)
	fmt.Printf("Port: %d\n", conf.Port)
	fmt.Printf("User: %s\n", conf.User)
	fmt.Printf("Pass: %s\n", conf.Pass)
	fmt.Printf("Db: %s\n", conf.Db)

	fmt.Println("====================================")

	fmt.Printf("Migration Path: %s\n", newGlobalConfig.Migration.Path)
	fmt.Printf("Database Host: %s\n", newGlobalConfig.Database.Host)
	fmt.Printf("Database User: %s\n", newGlobalConfig.Database.User)
	fmt.Printf("Database Pass: %s\n", newGlobalConfig.Database.Pass)
	fmt.Printf("Database DB: %s\n", newGlobalConfig.Database.Db)
	fmt.Printf("Deploy clients: %v\n", newGlobalConfig.Deploy.Clients)
	fmt.Printf("Servers: %v\n", newGlobalConfig.Servers)
	fmt.Printf("Server alpha: %v\n", newGlobalConfig.Servers["alpha"])
	fmt.Printf("Server alpha ip: %v\n", newGlobalConfig.Servers["alpha"].IP)
	fmt.Printf("Server alpha dc: %v\n", newGlobalConfig.Servers["alpha"].DC)
	fmt.Printf("Server beta: %v\n", newGlobalConfig.Servers["beta"])
	fmt.Printf("Server beta ip: %v\n", newGlobalConfig.Servers["beta"].IP)
	fmt.Printf("Server beta dc: %v\n", newGlobalConfig.Servers["beta"].DC)

}
