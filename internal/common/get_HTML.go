package common

import (
	"fmt"
	"io"
	"strings"
)

func (c *Client) GetHTML(rawUrl string) (string, error) {
	resp, err := c.httpClient.Get(rawUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error level status code: %d", resp.StatusCode)
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("content type is not text/html. Current content type:%s", resp.Header.Get("Content-Type"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
