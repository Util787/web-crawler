package common

import (
	"net/http"
	"time"
)

// better use this than create new http.Client every time
type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}
