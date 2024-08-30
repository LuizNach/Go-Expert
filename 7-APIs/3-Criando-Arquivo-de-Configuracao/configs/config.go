package configs

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JWTExpiresIn  int
}

var cfg *conf

// func LoadConfig(path string) (*conf, error) {
// 	// ...
// }
