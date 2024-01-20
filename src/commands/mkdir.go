package helpers

import (
	"os"
)

func MkDir(directoryName string) {
	if err := os.MkdirAll(directoryName, 0755); err != nil {
		ErrorHandle(err)
	}
}
