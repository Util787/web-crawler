package commands

import "fmt"

func Help() {
	fmt.Println("Available commands:")
	for _, usage := range cmdUsages {
		fmt.Printf("  %v\n", usage)
	}
}
