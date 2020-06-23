package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	file, err := os.Create("index.html")

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	h := []region{
		{
			Region: "Southern",
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
					Region:  "southern",
				},
			},
		},
		{
			Region: "Northern",
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
					Region:  "southern",
				},
			},
		},
		{
			Region: "Central",
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
					Region:  "southern",
				},
			},
		},
		{
			Region: "Central",
		},
	}

	fmt.Println(h)

	err = tpl.Execute(file, h)
	if err != nil {
		log.Fatalln(err)
	}
}
