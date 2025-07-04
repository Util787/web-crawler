package crawler

import (
	"log/slog"
	"sync"
	"time"

	"github.com/Util787/task-manager/pkg/logger/sl"
	"github.com/Util787/web-crawler/internal/common"
)

// TODO: add db
type Crawler struct {
	Log                *slog.Logger
	client             *common.Client
	BaseURL            string
	Pages              map[string]struct{}
	maxPages           int
	Wg                 *sync.WaitGroup
	concurrencyControl chan struct{}
	mu                 *sync.Mutex
}

func New(httpClientTimeout time.Duration, log *slog.Logger, baseURL string, concurrencyLimit int, maxPages int) *Crawler {
	pages := make(map[string]struct{})
	normalizedBaseURL, err := common.NormalizeURL(baseURL)
	if err != nil {
		log.Error("Error normalizing base URL", sl.Err(err))
		return nil
	}
	pages[normalizedBaseURL] = struct{}{}

	return &Crawler{
		BaseURL:            baseURL,
		client:             common.NewClientWithTimeout(httpClientTimeout),
		Log:                log,
		Pages:              pages,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, concurrencyLimit),
		Wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}
}
