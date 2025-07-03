package common

import (
	"fmt"
	"net/url"
	"strings"
)

func ValidateURL(rawUrl string) error {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}

	if parsedUrl.Scheme == "" {
		return fmt.Errorf("missing scheme (http:// or https://)")
	}
	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return fmt.Errorf("unsupported scheme: %s (http:// or https:// only)", parsedUrl.Scheme)
	}

	if parsedUrl.Host == "" {
		return fmt.Errorf("URL has no host")
	}
	if !strings.Contains(parsedUrl.Host, ".") {
		return fmt.Errorf("invalid host: %s", parsedUrl.Host)
	}
	return nil
}
