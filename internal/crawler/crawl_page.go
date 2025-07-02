package crawler

import (
	"log/slog"

	"github.com/Util787/task-manager/pkg/logger/sl"
	"github.com/Util787/web-crawler/internal/common"
)

func (c *Crawler) CrawlPage(baseUrl, currentUrl string, pages map[string]struct{}) {

	c.log.Info("Crawling page", slog.String("current_url", currentUrl))
	html, err := c.client.GetHTML(currentUrl)
	if err != nil {
		c.log.Error("Error getting HTML", sl.Err(err))
		return
	}

	urls, err := common.GetURLsFromHTML(html, baseUrl)
	if err != nil {
		c.log.Error("Error getting URLs", sl.Err(err))
		return
	}

	for _, url := range urls {
		err := common.ValidateURLDomain(baseUrl, url)
		if err != nil {
			c.log.Debug("Current URL is not on the same domain as the base URL", sl.Err(err), slog.String("current_url", url), slog.String("base_url", baseUrl))
			continue
		}

		normalizedUrl, err := common.NormalizeURL(url)
		if err != nil {
			c.log.Error("Error normalizing URL", sl.Err(err))
			continue
		}

		if _, ok := pages[normalizedUrl]; ok {
			continue
		}
		pages[normalizedUrl] = struct{}{}

		c.CrawlPage(baseUrl, url, pages)
	}
}
