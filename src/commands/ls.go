package helpers

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func Ls() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)

	entries, err := os.ReadDir("./")
	ErrorHandle(err)

	fmt.Fprintln(w, "Name\tSize\tIs Directory")
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Println("Error getting info for entry:", entry.Name())
			continue
		}
		fmt.Fprintf(w, "%s\t%d\t%v\n", entry.Name(), info.Size(), entry.IsDir())
	}
	w.Flush()
}
