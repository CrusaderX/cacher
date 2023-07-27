package fetcher

type Fetcher interface {
	Name() string
	Fetch() []string
}
