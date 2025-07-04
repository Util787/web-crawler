package commands

import (
	"fmt"

	"github.com/Util787/web-crawler/internal/crawler"
)

var ShowParamsUsage = fmt.Sprintf("%v — show current parameters", ShowParamsCommand)

func ShowParams(c *crawler.Crawler, httpClientTimeout int, concurrencyLimit int) {
	fmt.Println("HTTP Client Timeout:", httpClientTimeout)
	fmt.Println("Concurrency Limit:", concurrencyLimit)
	fmt.Println("Base URL:", c.BaseURL)
}
