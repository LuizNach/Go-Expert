package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"` // o viper sabe estruturar o campo quevier DB_DRIVER no campo DBDriver se usarmos a annotation
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

var cfg *conf

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config") // Cria o nome da config que podemos recuperar futuramente
	viper.SetConfigType("env")        // Tipo de carregamento que vamos utilizar: env, json, toml etc
	viper.AddConfigPath(path)         // Caminho para encontrar o arquivo responsavel pelas env vars
	viper.SetConfigFile(".env")       // Nome do arquivo buscado no path

	viper.AutomaticEnv() // Carrega env vars já existentes no ambiente.
	// Tenta carregar as env vars que já estão no map do viper pelas que existirem no ambiente.
	// Mesmo que vc tenha caregado um arquivo .env o AutomaticEnv vai pegar das variaveis de ambiente.
	// Assim substituimos o que carregamos no .env por variáveis que já existem no ambiente.

	err := viper.ReadInConfig()
	if err != nil {
		// return nil, err
		panic(err)
	}

	err = viper.Unmarshal(&cfg) // unmarshals the config into a Struct.
	if err != nil {
		// return nil, err
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
