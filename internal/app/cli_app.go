package app

import (
	"fmt"
	"os"

	"github.com/Util787/web-crawler/internal/common"
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

	client := common.NewClient()
	html, err := client.GetHTML(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	fmt.Println(html)
}
