package helpers

import (
	"log"
	"os"
)

func CreateFile(fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
