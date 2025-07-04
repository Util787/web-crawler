package commands

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Util787/web-crawler/internal/crawler"
)

var CrawlPageUsage = fmt.Sprintf("%v <url> — crawl a page (if url is not provided, will crawl baseURL)", CrawlCommand)

func CrawlPage(c *crawler.Crawler, url string) {
	start := time.Now()
	c.Log.Info("Starting crawl", slog.Any("time", start))
	defer func() {
		c.Log.Info("Crawl ended", slog.Any("time", time.Now()), slog.Duration("duration", time.Duration(time.Since(start).Milliseconds())))
	}()

	c.Wg.Add(1)
	go func() {
		defer c.Wg.Done()
		c.CrawlPage(url)
	}()
	c.Wg.Wait()
}
