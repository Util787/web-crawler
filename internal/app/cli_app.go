package app

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Util787/web-crawler/internal/crawler"
)

const httpClientTimeout = 5 * time.Second
const concurrencyLimit = 100

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

	baseURL := args[1]
	log.Info("Starting crawler", slog.String("url", baseURL))

	c := crawler.New(httpClientTimeout, log, baseURL, concurrencyLimit)

	start := time.Now()
	defer func() {
		log.Info("Crawler took", slog.Duration("duration", time.Duration(time.Since(start).Milliseconds())))
	}()

	c.Wg.Add(1)
	go func() {
		defer c.Wg.Done()
		c.CrawlPage(baseURL)
	}()
	c.Wg.Wait()

	log.Info("Found pages after crawl", slog.Int("pages_length", len(c.Pages)), slog.Any("pages", c.Pages))
}

// for test: go run cmd/main.go https://www.wagslane.dev
