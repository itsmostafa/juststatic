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

	// Context variables
	Context struct {
		Title  string `json:"title"`
		Slogan string `json:"slogan"`
		Phone  string `json:"phone"`
		Email  string `json:"email"`
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
		data, _ := ioutil.ReadFile("data.json")
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
