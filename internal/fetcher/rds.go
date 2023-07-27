package fetcher

type Rds struct {
	name string
	tag  string
}

func NewRds(name, tag string) *Rds {
	return &Rds{
		name: name,
		tag:  tag,
	}
}

func (e *Rds) Name() string {
	return e.name
}

func (e *Rds) Fetch() []string {
	return []string{"3", "4"}
}
