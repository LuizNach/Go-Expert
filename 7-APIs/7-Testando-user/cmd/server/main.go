package main

import (
	"github.com/LuizNach/Go-Expert/7-APIs/7-Testando-user/configs"
)

func main() {

	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)
}
