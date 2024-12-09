package main

import (
	"github.com/LuizNach/Go-Expert/7-APIs/6-Criando-Entidade-User/configs"
)

func main() {

	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)
}
