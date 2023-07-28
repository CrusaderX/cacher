package fetcher

type Fetcher interface {
	Name() string
	Fetch() *[]Resource
}

type Tags struct {
	Name string
}

type Resource struct {
	Name *string
	Tags *Tags
}
