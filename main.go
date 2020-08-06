package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Project struct {
	Name        string `toml:"name"`
	Language    string `toml:"language"`
	Tag         string `toml:"tag"`
	Description string `toml:"description"`
	Type        string `toml:"type"`
}

type Projects struct {
	P []Project `toml:"project"`
}

func main() {
	var projects Projects
	tomlData, _ := ioutil.ReadFile("projects.toml")
	_, err := toml.Decode(string(tomlData), &projects)
	if err != nil {
		panic(err)
	}
	fmt.Println(projects)
	fmt.Println("vim-go")
}
