package commands

import "bufio"

// set params should only be used on first input, nil pointers are used to get new values instead of keeping current values
func SetParams(reader *bufio.Reader) (httpClientTimeout int, concurrencyLimit int, baseURL string, maxPages int) {
	httpClientTimeout = getHttpClientTimeout(reader, nil)
	concurrencyLimit = getConcurrencyLimit(reader, nil)
	baseURL = getBaseURL(reader, nil)
	maxPages = getMaxPages(reader, nil)
	return httpClientTimeout, concurrencyLimit, baseURL, maxPages
}
