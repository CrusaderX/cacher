package registry

import (
	"sync"
	"time"

	"github.com/CrusaderX/cacher/internal/fetcher"
)

type FetcherRegistry struct {
	fetchers map[string]fetcher.Fetcher
	results  chan Result
	period   time.Duration
}

func NewFetcherRegistry(period time.Duration) *FetcherRegistry {
	return &FetcherRegistry{
		fetchers: make(map[string]fetcher.Fetcher),
		results:  make(chan Result),
		period:   period,
	}
}

func (r *FetcherRegistry) Close() {
	close(r.results)
}

func (r *FetcherRegistry) Results() <-chan Result {
	return r.results
}

func (r *FetcherRegistry) Register(fetcher fetcher.Fetcher) {
	r.fetchers[fetcher.Name()] = fetcher
}

func (r *FetcherRegistry) Fetch() {
	wg := sync.WaitGroup{}

	for _ = range time.Tick(r.period) {
		for _, fth := range r.fetchers {
			wg.Add(1)

			go func(f fetcher.Fetcher) {

				defer wg.Done()

				values := f.Fetch()
				r.results <- Result{
					FetcherID: f.Name(),
					Values:    values,
				}
			}(fth)
		}

		wg.Wait()
	}
}
