package crawler

import (
	"log/slog"

	"github.com/Util787/task-manager/pkg/logger/sl"
	"github.com/Util787/web-crawler/internal/common"
)

func (c *Crawler) CrawlPage(currentUrl string) {

	c.concurrencyControl <- struct{}{}
	defer func() { <-c.concurrencyControl }()

	c.Log.Info("Crawling page", slog.String("current_url", currentUrl))
	html, err := c.client.GetHTML(currentUrl)
	if err != nil {
		c.Log.Error("Error getting HTML", sl.Err(err))
		return
	}

	urls, err := common.GetURLsFromHTML(html, c.BaseURL)
	if err != nil {
		c.Log.Error("Error getting URLs", sl.Err(err))
		return
	}
	c.Log.Debug("Found URLs", slog.Int("urls_length", len(urls)), slog.Any("urls", urls), slog.String("current_url", currentUrl))

	for _, url := range urls {
		err := common.ValidateURLDomain(c.BaseURL, url)
		if err != nil {
			continue
		}

		normalizedUrl, err := common.NormalizeURL(url)
		if err != nil {
			c.Log.Error("Error normalizing URL", sl.Err(err))
			continue
		}

		c.mu.Lock()
		if c.maxPages > 0 && len(c.Pages) >= c.maxPages {
			c.Log.Info("Max pages reached", slog.Int("max_pages", c.maxPages), slog.Int("current_pages_length", len(c.Pages)))
			c.mu.Unlock()
			return
		}
		if _, ok := c.Pages[normalizedUrl]; ok {
			c.mu.Unlock()
			continue
		}
		c.Pages[normalizedUrl] = struct{}{}
		c.mu.Unlock()

		c.Wg.Add(1)
		go func() {
			defer c.Wg.Done()
			c.CrawlPage(url)
		}()

	}
}
