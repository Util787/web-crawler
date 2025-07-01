package crawler

import (
	"log/slog"

	"github.com/Util787/task-manager/pkg/logger/sl"
	"github.com/Util787/web-crawler/internal/common"
)

func (c *Crawler) CrawlPage(baseUrl, currentUrl string, pages map[string]int) error {

	err := common.ValidateURLDomain(baseUrl, currentUrl)
	if err != nil {
		c.log.Error("Error validating URL's domain", sl.Err(err), slog.String("current url", currentUrl), slog.String("base url", baseUrl))
		return err
	}

	c.log.Info("Crawling page", slog.String("url", currentUrl))
	html, err := c.client.GetHTML(currentUrl)
	if err != nil {
		c.log.Error("Error getting HTML", sl.Err(err))
		return err
	}

	urls, err := common.GetURLsFromHTML(html, baseUrl)
	if err != nil {
		c.log.Error("Error getting URLs", sl.Err(err))
		return err
	}

	for _, url := range urls {
		if _, ok := pages[url]; ok {
			continue
		}

		normalizedUrl, err := common.NormalizeURL(url)
		if err != nil {
			c.log.Error("Error normalizing URL", sl.Err(err))
			continue
		}
		pages[normalizedUrl] = 1

		if err := c.CrawlPage(baseUrl, normalizedUrl, pages); err != nil {
			c.log.Error("Error crawling page", sl.Err(err))
		}
	}

	return nil
}
