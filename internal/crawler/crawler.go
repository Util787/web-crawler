package crawler

import (
	"log/slog"
	"time"

	"github.com/Util787/web-crawler/internal/common"
)

// TODO: add db, make output to file
type Crawler struct {
	client *common.Client
	log    *slog.Logger
}

func New(httpClientTimeout time.Duration, log *slog.Logger) *Crawler {
	return &Crawler{client: common.NewClientWithTimeout(httpClientTimeout), log: log}
}
