package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"
)

// Page ...
type Page struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// generate ...
func generate() error {
	file, _ := ioutil.ReadFile("data.json")
	page := Page{}

	_ = json.Unmarshal([]byte(file), &page)

	distFile, err := os.Create("public/index.html")
	if err != nil {
		return err
	}
	defer distFile.Close()

	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(distFile, page)
	return nil
}

func main() {
	generate()
}
