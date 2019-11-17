package main

import (
	"fmt"
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Project struct {
		Name        string
		Description string
	}
	File []struct {
		Path        string
		Description string
		Text        string
	}
}

func main() {
	fmt.Println("hello")
	f, err := os.Open("conf/sample.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var config Config
	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
	}

	for _, v := range config.File {
		fmt.Println(v.Path)
	}

}
