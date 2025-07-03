package crawler

import (
	"log/slog"
	"sync"
	"time"

	"github.com/Util787/task-manager/pkg/logger/sl"
	"github.com/Util787/web-crawler/internal/common"
)

// TODO: add db, make output to file
type Crawler struct {
	Pages              map[string]struct{}
	client             *common.Client
	log                *slog.Logger
	baseURL            string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func New(httpClientTimeout time.Duration, log *slog.Logger, baseURL string) *Crawler {
	pages := make(map[string]struct{})
	normalizedBaseURL, err := common.NormalizeURL(baseURL)
	if err != nil {
		log.Error("Error normalizing base URL", sl.Err(err))
		return nil
	}
	pages[normalizedBaseURL] = struct{}{}

	return &Crawler{
		baseURL:            baseURL,
		client:             common.NewClientWithTimeout(httpClientTimeout),
		log:                log,
		Pages:              pages,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
	}
}
