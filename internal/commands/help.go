package commands

import "fmt"

func Help() {
	fmt.Println("Available commands:")
	fmt.Printf("  %v <url> — crawl a page (if url is not provided, will crawl baseURL)\n", CrawlCommand)
	fmt.Printf("  %v — show current parameters\n", ShowParamsCommand)
	fmt.Printf("  %v — change parameters (timeout, concurrency, baseURL)\n", ResetParamsCommand)
	fmt.Printf("  %v — show this help message\n", HelpCommand)
	fmt.Printf("  %v — exit the program\n", ExitCommand)
}
