package app

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Util787/web-crawler/internal/crawler"
)

const httpClientTimeout = 5 * time.Second

func Run(log *slog.Logger) {
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
	log.Info("Starting crawler", slog.String("url", url))

	c := crawler.New(httpClientTimeout, log)
	pages := make(map[string]struct{})
	c.CrawlPage(url, url, pages)

	log.Info("Found pages after crawl", slog.Int("pages_length", len(pages)), slog.Any("pages", pages))
}
