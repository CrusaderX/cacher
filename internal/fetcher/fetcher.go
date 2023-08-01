package fetcher

type Fetcher interface {
	Name() string
	Fetch() *[]Resource
}

type Resource struct {
	Namespace map[string][]string `json:"namespace"`
}
