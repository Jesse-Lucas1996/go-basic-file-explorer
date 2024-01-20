package helpers

import (
	"log"
	"os"
)

func Cd(directoryChange string) {
	entries, err := os.ReadDir(directoryChange)
	ErrorHandle(err)
	err = os.Chdir(directoryChange)
	ErrorHandle(err)
	for _, entry := range entries {
		log.Println(entry.Name())
	}
}
