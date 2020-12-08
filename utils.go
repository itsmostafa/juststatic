package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

// File channel initialize
type (
	File chan string

	// Context for templates
	Context struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Domain      string `json:"domain"`
		Phone       string `json:"phone"`
		Address     string `json:"address"`
		Email       string `json:"email"`
	}
)

// Walk through folder
func (f File) Walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		f <- path
	}

	return nil
}

// Parse templates
func Parse(src, dest, file string) error {
	if file != "base.html" {
		data, _ := ioutil.ReadFile(filepath.Join(src, "data.json"))
		context := Context{}

		_ = json.Unmarshal([]byte(data), &context)

		distFile, err := os.Create(filepath.Join(dest, file))
		if err != nil {
			return err
		}
		defer distFile.Close()

		t := template.Must(template.ParseFiles(filepath.Join(src, "base.html"), filepath.Join(src, file)))
		err = t.ExecuteTemplate(distFile, "base", context)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
