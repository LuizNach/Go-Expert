package main

type Config2 struct {
	Migration MigrationConfig         `toml:"migration"`
	Database  DatabaseConfig          `toml:"database"`
	Deploy    DeployConfig            `toml:"deploy"`
	Servers   map[string]ServerConfig `toml:"servers"`
}

type MigrationConfig struct {
	Path string `toml:"path"`
}

type DatabaseConfig struct {
	Host string `toml:"host"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
	Db   string `toml:"db"`
}

type DeployConfig struct {
	Clients []int `toml:"clients"`
}

type ServerConfig struct {
	IP string `toml:"ip"`
	DC string `toml:"dc"`
}
