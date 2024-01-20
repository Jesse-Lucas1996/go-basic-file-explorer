package helpers

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func Search(fileName string) {
	found := false
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}
		if d.Name() == fileName {
			found = true
			if d.IsDir() {
				log.Println("Found directory!", d.Name())
			} else {
				log.Println("Found file!", d.Name())
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if !found {
		dir, _ := os.Getwd()
		log.Println("Not found in current directory:", dir)
	}
}
