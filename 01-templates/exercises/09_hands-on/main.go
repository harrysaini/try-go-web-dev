package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

// Record of data
type Record struct {
	Date time.Time
	Open float64
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func checkDefer() {
	fmt.Println("deferred")
}

func createFile() *os.File {
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	defer checkDefer()

	return file
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":3005", nil)
}

func handleRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)

	records := parseCSVFile()
	tmpl.Execute(res, records)
}

func parseCSVFile() (records []Record) {
	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records = make([]Record, 0, len(rows))

	for _, row := range rows[1:] {
		date, _ := time.Parse("2006-01-02", row[0])
		round, _ := strconv.ParseFloat(row[1], 64)
		r := Record{
			Date: date,
			Open: round,
		}
		records = append(records, r)
	}

	return records
}
