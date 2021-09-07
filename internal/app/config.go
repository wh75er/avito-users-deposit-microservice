package app

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	Server struct {
		Port int
	}
	Storage struct {
		Host string
		Port int
		MaxPoolConn int
	}
}

func (c *config) loadFromToml(tomlData string) {
	if _, err := toml.DecodeFile("configs/" + tomlData, c); err != nil {
		log.Fatal("Failed to decode toml file: " + err.Error())
	}
	fmt.Println(c)
}
