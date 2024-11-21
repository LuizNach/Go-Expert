package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

// Reference: https://aprendagolang.com.br/carregando-configuracoes-com-viper/

type MysqlCfg struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func init() {
	viper.SetConfigName("config")    // nome do arquivo que queremos carregar
	viper.SetConfigType("json")      //extensão do arquivo
	viper.AddConfigPath(".")         // caminho onde o arquivo está
	viper.AddConfigPath("./config/") // caminho alternativo onde o arquivo está
	err := viper.ReadInConfig()      // lê o arquivo e carrega seu conteúdo
	if err != nil {
		panic(fmt.Errorf("erro fatal no arquivo de configuração: %w", err))
	}

}

func main() {
	fmt.Printf("Pegando uma string dado uma chave do json, \"environment\": %v\n", viper.GetString("environment"))

	// Para recuperar valores que estão encadeados, tudo que precisamos fazer é chamar todas as chaves usando um
	// entre elas, como podemos ver no exemplo abaixo, onde vamos printar o host do mysql

	fmt.Printf("Pegando a string nested dentro do json, \"databases.mysql.host\": %v\n", viper.GetString("databases.mysql.host"))

	bytes, err := json.Marshal(viper.Get("databases.mysql"))
	if err != nil {
		panic(err)
	}

	var mysql MysqlCfg
	err = json.Unmarshal(bytes, &mysql)
	if err != nil {
		panic(err)
	}

	fmt.Printf("A struct MysqlCfg: %v\n", mysql)
	fmt.Printf("MysqlCfg Host: %v\n", mysql.Host)
	fmt.Printf("MysqlCfg Port: %v\n", mysql.Port)
	fmt.Printf("MysqlCfg User: %v\n", mysql.User)
	fmt.Printf("MysqlCfg Pass: %v\n", mysql.Pass)

	// Aqui tem um bom exemplo da keyword SetDefault para caso valores não existam
	viper.SetDefault("databases.mongodb.host", "localhost:3000") // caso nao exista vai pegar essa aqui
	viper.SetDefault("chave_que_nao_existe", "blahblah")
	fmt.Printf("Pegando databases.mongodb.host que ja existe: %v\n", viper.GetString("databases.mongodb.host"))
	fmt.Printf("Pegando key que nao existe no json: %v\n", viper.GetString("chave_que_nao_existe"))

}
