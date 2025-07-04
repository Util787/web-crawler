package common

import (
	"fmt"
	"net/url"
	"strings"
)

// If we have the following urls:
//
// https://someurl.com/path/
//
// https://someurl.com/path
//
// http://someurl.com/path/
//
// http://someurl.com/path
//
// NormalizeUrl will normalize those urls to the someurl.com/path
func NormalizeURL(rawUrl string) (string, error) {
	if strings.TrimSpace(rawUrl) == "" {
		return "", fmt.Errorf("empty raw url")
	}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	// If the scheme is invalid then url.Parse wont work correctly
	if err := ValidateURL(rawUrl); err != nil {
		return "", err
	}

	normalized := parsedUrl.Host + strings.TrimRight(parsedUrl.Path, "/")
	return normalized, nil
}
