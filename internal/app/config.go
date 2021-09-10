package app

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	Server Server
	Storage Storage
}

type Server struct {
	Port int
}

type Storage struct {
	Url string
	Driver string
	MaxPoolConn int
}

func newConfig() *config {
	return &config {
		Server{
			3000,
		},
		Storage {
			"postgresql://postgres:postgres@localhost:5432/postgres",
			"pgx",
			30,
		},
	}
}

func (c *config) loadFromToml(tomlData string) (e error) {
	_, e = toml.DecodeFile("configs/" + tomlData, c)
	return
}
