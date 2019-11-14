package main

import (
	"fmt"
	_ "github.com/BurntSushi/toml"
	_ "github.com/naoina/toml"
	"github.com/pelletier/go-toml"
)

type Config struct {
	Path map[string]PathConfig
}

type PathConfig struct {
	Description string
	Text        string
}

func main() {
	fmt.Println("hello")
	config, err := toml.LoadFile("sample_conf/config.toml")
	if err != nil {
		panic(err)
	}

	fmt.Println(config.ToTomlString())

}
