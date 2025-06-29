package common

import (
	"net/url"
	"strings"
)

//If we have the following urls:
// https://someurl.com/path/
// https://someurl.com/path
// http://someurl.com/path/
// http://someurl.com/path

// NormalizeUrl will normalize those urls to the someurl.com/path
func NormalizeUrl(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	normalized := parsedUrl.Host + strings.TrimRight(parsedUrl.Path, "/")
	return normalized, nil
}
