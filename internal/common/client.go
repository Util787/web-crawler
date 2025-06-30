package common

import (
	"net/http"
	"time"
)

// better use this than create new http.Client every time
type Client struct {
	httpClient *http.Client
}

func NewClientWithTimeout(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
	}
}
