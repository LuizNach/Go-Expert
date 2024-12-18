package main

import (
	"github.com/LuizNach/Go-Expert/7-APIs/9-Testando-Product/configs"
)

func main() {

	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)
}
