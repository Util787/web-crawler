package crawler

import (
	"time"

	"github.com/Util787/web-crawler/internal/common"
)

// TODO: add db
type Crawler struct {
	client *common.Client
}

func New(httpClientTimeout time.Duration) *Crawler {
	return &Crawler{client: common.NewClientWithTimeout(httpClientTimeout)}
}
