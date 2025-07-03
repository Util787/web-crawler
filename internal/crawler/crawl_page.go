package crawler

import (
	"log/slog"

	"github.com/Util787/task-manager/pkg/logger/sl"
	"github.com/Util787/web-crawler/internal/common"
)

func (c *Crawler) CrawlPage(currentUrl string) {

	c.log.Info("Crawling page", slog.String("current_url", currentUrl))
	html, err := c.client.GetHTML(currentUrl)
	if err != nil {
		c.log.Error("Error getting HTML", sl.Err(err))
		return
	}

	urls, err := common.GetURLsFromHTML(html, c.baseURL)
	if err != nil {
		c.log.Error("Error getting URLs", sl.Err(err))
		return
	}
	c.log.Debug("Found URLs", slog.Int("urls_length", len(urls)), slog.Any("urls", urls), slog.String("current_url", currentUrl))

	for _, url := range urls {
		err := common.ValidateURLDomain(c.baseURL, url)
		if err != nil {
			c.log.Debug("Current URL is not on the same domain as the base URL", sl.Err(err), slog.String("current_url", url), slog.String("base_url", c.baseURL))
			continue
		}

		normalizedUrl, err := common.NormalizeURL(url)
		if err != nil {
			c.log.Error("Error normalizing URL", sl.Err(err))
			continue
		}

		if _, ok := c.Pages[normalizedUrl]; ok {
			continue
		}
		c.Pages[normalizedUrl] = struct{}{}

		c.CrawlPage(url)
	}
}
