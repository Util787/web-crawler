package crawler

import (
	"time"

	"github.com/Util787/web-crawler/internal/common"
)

// TODO: add db
type Crawler struct {
	Client *common.Client
}

func New(httpClientTimeout time.Duration) *Crawler {
	return &Crawler{Client: common.NewClientWithTimeout(httpClientTimeout)}
}
