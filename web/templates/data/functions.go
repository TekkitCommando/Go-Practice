package main

import (
	"text/template"
	"strings"
	"os"
	"log"
)

type sage struct {
	Name string
	Motto string
}
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.New("tpl4.gohtml").Funcs(fm).ParseFiles("tpl4.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	b := sage{
		Name: "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage {
		Name: "Gandhi",
		Motto: "Be the change",
	}

	m := sage {
		Name: "MLK",
		Motto: "Some motto",
	}

	sages := []sage{b, g, m}

	data := struct {
		Wisdom []sage
	}{
		sages,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}