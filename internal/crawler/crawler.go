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
	Pages              map[string]struct{}
	client             *common.Client
	Log                *slog.Logger
	BaseURL            string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
}

func New(httpClientTimeout time.Duration, log *slog.Logger, baseURL string, concurrencyLimit int) *Crawler {
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
	}
}
