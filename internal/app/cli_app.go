package app

import (
	"fmt"
	"os"
)

func Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(args)
		fmt.Println("Not enough arguments. Usage: go run cmd/main.go <url>")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println(args)
		fmt.Println("Too many arguments. Usage: go run cmd/main.go <url>")
		os.Exit(1)
	}

	url := args[1]
	fmt.Printf("Starting crawl of:%s\n", url)
}
