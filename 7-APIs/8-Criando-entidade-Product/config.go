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

func LoadConfig(path string) (*conf, error) {
	var cfg *conf

	viper.SetConfigName("app_config") // Cria o nome da config que podemos recuperar futuramente
	viper.SetConfigType("env")        // Tipo de carregamento que vamos utilizar: env, json, toml etc
	viper.AddConfigPath(path)         // Caminho para encontrar o arquivo responsavel pelas env vars
	viper.SetConfigFile(".env")       // Nome do arquivo buscado no path

	viper.ReadInConfig() // Carrega as variaveis encontradas no viper

	viper.SetDefault("DB_DRIVER", "postgres") // estabelecemos o valor padrao caso a gente não encontre a variavel DB_DRIVER
	viper.SetDefault("DB_HOST", "LOCALHOST")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "root")
	viper.SetDefault("DB_NAME", "fullcycle")
	viper.SetDefault("WEB_SERVER_PORT", "8000")
	viper.SetDefault("JWT_SECRET", "secret")
	viper.SetDefault("JWT_EXPIRES_IN", "300")

	cfg = &conf{
		DBDriver:      viper.GetString("DB_DRIVER"), // Buscamos na env var o  valor de DB_DRIVER
		DBHost:        viper.GetString("DB_HOST"),
		DBPort:        viper.GetString("DB_PORT"),
		DBUser:        viper.GetString("DB_USER"),
		DBPassword:    viper.GetString("DB_PASSWORD"),
		DBName:        viper.GetString("DB_NAME"),
		WebServerPort: viper.GetString("WEB_SERVER_PORT"),
		JWTSecret:     viper.GetString("JWT_SECRET"),
		JWTExpiresIn:  viper.GetInt("JWT_EXPIRES_IN"),
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
