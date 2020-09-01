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
	f.WriteString(`
# Repo Index

These are not all my projects.
These are just the ` + fmt.Sprint(len(projects.P)) + ` projects that I enjoy the most and actively work to improve.
`)

	// tag -> type -> project
	p := make(map[string]map[string][]Project)
	p2 := make(map[string][]Project)
	for _, project := range projects.P {
		if _, ok := p[project.Tag]; !ok {
			p[project.Tag] = make(map[string][]Project)
			p[project.Tag][project.Type] = []Project{}
		}
		if _, ok := p2[project.Type]; !ok {
			p2[project.Type] = []Project{}
		}
		p2[project.Type] = append(p2[project.Type], project)
		p[project.Tag][project.Type] = append(p[project.Tag][project.Type], project)
	}

	f.WriteString(`
### organized by topic

`)

	for k := range p {
		i := 0
		num := 0
		for l := range p[k] {
			num += len(p[k][l])
		}

		for l := range p[k] {
			// if len(p[k][l]) == 1 {
			// 	continue
			// }
			if i == 0 {
				i++
				f.WriteString(fmt.Sprintf("<details><summary><strong style='font-size:2rem;'>%s (%d projects)</strong></summary>\n\n", k, num))
			}

			f.WriteString(fmt.Sprintf("<h3>%s</h3><ul>\n\n", l))
			for _, q := range p[k][l] {
				f.WriteString(fmt.Sprintf("<li><a href='https://github.com/schollz/%s'>%s</a>: %s (%s)</li>\n\n", q.Name, q.Name, q.Description, q.Language))
			}
			f.WriteString("</ul>")
		}
		f.WriteString("\n\n</details>\n")
	}
	// f.WriteString(fmt.Sprintf("## other\n\n"))
	// for k := range p {
	// 	for l := range p[k] {
	// 		if len(p[k][l]) > 1 {
	// 			continue
	// 		}
	// 		for _, q := range p[k][l] {
	// 			f.WriteString(fmt.Sprintf("[%s](https://github.com/schollz/%s): %s\n\n", q.Name, q.Name, q.Description))
	// 		}
	// 	}
	// }

	f.WriteString(`
### organized by type

`)
	for k := range p2 {
		num := len(p2[k])
		f.WriteString(fmt.Sprintf("<details><summary><strong style='font-size:2rem;'>%s (%d projects)</strong></summary>\n\n", k, num))
		f.WriteString(fmt.Sprintf("<ul>\n\n"))

		for _, q := range p2[k] {
			f.WriteString(fmt.Sprintf("<li><a href='https://github.com/schollz/%s'>%s</a>: %s (%s)</li>\n\n", q.Name, q.Name, q.Description, q.Language))
		}
		f.WriteString("</ul>")
		f.WriteString("\n\n</details>\n")
	}
}
