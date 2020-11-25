package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Page ...
type Page struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// copy an entire directory
func copy(source, destination string) error {
	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		relPath := strings.Replace(path, source, "", 1)

		switch ext := strings.ToLower(filepath.Ext(relPath)); ext {
		case ".html":
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
		default:
			if relPath == "" {
				return nil
			}
			if info.IsDir() {
				return os.Mkdir(filepath.Join(destination, relPath), 0755)
			}
			data, err := ioutil.ReadFile(filepath.Join(source, relPath))
			if err != nil {
				return err
			}
			return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
		}

	})
	return err
}

func main() {
	err := copy("templates", "public")
	if err != nil {
		fmt.Println(err)
	}
}
