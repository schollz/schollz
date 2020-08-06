package main

import (
	"fmt"
	"io/ioutil"
	"os"

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

	f, _ := os.Create("README.md")
	defer f.Close()

	// tag -> type -> project
	p := make(map[string]map[string][]Project)
	for _, project := range projects.P {
		if _, ok := p[project.Tag]; !ok {
			p[project.Tag] = make(map[string][]Project)
			p[project.Tag][project.Type] = []Project{}
		}
		p[project.Tag][project.Type] = append(p[project.Tag][project.Type], project)
	}

	for k := range p {
		i := 0
		for l := range p[k] {
			if len(p[k][l]) == 1 {
				continue
			}
			if i == 0 {
				i++
				f.WriteString(fmt.Sprintf("## %s\n\n", k))
			}

			f.WriteString(fmt.Sprintf("### %s\n\n", l))
			for _, q := range p[k][l] {
				f.WriteString(fmt.Sprintf("[%s](https://github.com/schollz/%s): %s\n\n", q.Name, q.Name, q.Description))
			}
		}
	}
	f.WriteString(fmt.Sprintf("## other\n\n"))
	for k := range p {
		for l := range p[k] {
			if len(p[k][l]) > 1 {
				continue
			}
			for _, q := range p[k][l] {
				f.WriteString(fmt.Sprintf("[%s](https://github.com/schollz/%s): %s\n\n", q.Name, q.Name, q.Description))
			}
		}
	}
}
