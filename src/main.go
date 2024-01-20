package main

import (
	"bufio"
	"fmt"
	helpers "go-file-explorer/src/commands"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the file explorer!")
	fmt.Println("Type help to see the available commands.")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("⚡️ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "help":
			fmt.Println("Commands available are: mkdir, create-file, search ls, clear, cd, exit")
		case "mkdir":
			helpers.MkDir(args[1])
		case "ls":
			helpers.Ls()
		case "cd":
			helpers.Cd(args[1])
		case "create-file":
			helpers.CreateFile(args[1])
		case "search":
			helpers.Search(args[1])
		case "clear":
			helpers.CallClear()
		case "exit":
			fmt.Println("Bye!")
			return
		}
	}
}
