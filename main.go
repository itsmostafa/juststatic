package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Route folders
type Route struct {
	source, destination string
}

// Generate static files
func (r *Route) Generate() error {
	walker := make(File)

	go func() {
		// Gather the files to upload by walking the path recursively
		if err := filepath.Walk(r.source, walker.Walk); err != nil {
			log.Fatalln("File walk failed:", err)
		}
		close(walker)
	}()

	for path := range walker {
		relPath, err := filepath.Rel(r.source, path)
		if err != nil {
			fmt.Println(err)
		}

		// create destination folder if it does not exist
		if _, err := os.Stat(filepath.Join(r.destination, relPath)); os.IsNotExist(err) {
			os.MkdirAll(filepath.Dir(filepath.Join(r.destination, relPath)), 0755)
		}
		switch ext := strings.ToLower(filepath.Ext(relPath)); ext {
		case ".html":
			err := Parse(r.source, r.destination, relPath)
			if err != nil {
				fmt.Println(err)
			}
		default:
			// initialize destination file
			destFile, _ := os.Create(filepath.Join(r.destination, relPath))
			defer destFile.Close()
			// read file from source
			data, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
			}
			// write file
			destFile.Write(data)
		}

		fmt.Println("created " + path)
	}
	return nil
}

func main() {
	var templateDir string
	flag.StringVar(&templateDir, "t", "./template", "Specify directory from which to read the templates")
	flag.Parse()
	route := Route{"templates/" + templateDir, "public"}

	err := route.Generate()
	if err != nil {
		fmt.Println(err)
	}
}
