package v1

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/Util787/web-crawler/internal/commands"
	"github.com/Util787/web-crawler/internal/crawler"
)

func Run(log *slog.Logger) {
	args := os.Args
	if len(args) < 4 {
		fmt.Println(args)
		fmt.Println("Not enough arguments. Usage: go run cmd/main.go <url> <http_client_timeout> <max_concurrency>")
		os.Exit(1)
	}
	if len(args) > 4 {
		fmt.Println(args)
		fmt.Println("Too many arguments. Usage: go run cmd/main.go <url> <http_client_timeout> <max_concurrency>")
		os.Exit(1)
	}

	baseURL := args[1]
	httpClientTimeout, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid HTTP client timeout. Usage: go run cmd/main.go <url> <http_client_timeout> <max_concurrency>")
		return
	}

	concurrencyLimit, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Invalid max concurrency. Usage: go run cmd/main.go <url> <http_client_timeout> <max_concurrency>")
		return
	}

	log.Info("Starting crawler", slog.String("url", baseURL))

	c := crawler.New(time.Second*time.Duration(httpClientTimeout), log, baseURL, concurrencyLimit)

	commands.CrawlPage(c, baseURL)

	log.Info("Found pages after crawl", slog.Int("pages_length", len(c.Pages)), slog.Any("pages", c.Pages))

}

// for test: go run cmd/main.go https://www.wagslane.dev 5 100
