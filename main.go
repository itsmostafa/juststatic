package main

import (
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
	// create destination folder if it does not exist
	if _, err := os.Stat(r.destination); os.IsNotExist(err) {
		os.Mkdir(r.destination, 0755)
	}

	walker := make(File)

	go func() {
		// Gather the files to upload by walking the path recursively
		if err := filepath.Walk(r.source, walker.Walk); err != nil {
			log.Fatalln("File walk failed:", err)
		}
		close(walker)
	}()

	for path := range walker {
		relPath := strings.Replace(path, r.source+"/", "", 1)
		switch ext := strings.ToLower(filepath.Ext(relPath)); ext {
		case ".html":
			err := Parse(r.source, r.destination, relPath)
			if err != nil {
				fmt.Println(err)
			}
		default:
			// Read file from source
			data, _ := ioutil.ReadFile(path)
			// Write file to destination
			err := ioutil.WriteFile(filepath.Join(r.destination, relPath), data, 0777)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("created " + path)
	}
	return nil
}

func main() {
	route := Route{"template", "public"}

	err := route.Generate()
	if err != nil {
		fmt.Println(err)
	}
}
