package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tmpl *template.Template

var fns template.FuncMap = template.FuncMap{
	"uc": strings.ToUpper,
	"test": func(val ...string) string {
		// fmt.Println(val)
		return strings.Join(val, "-")
	},
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fns).ParseGlob("templates/*"))
}

func main() {
	names := []string{"ram", "shyam", "seeta"}

	err := tmpl.ExecuteTemplate(os.Stdout, "page.gohtml", names)

	if err != nil {
		log.Fatalln(err)
	}

}
