package fetcher

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
)

type RdsFilter struct {
	StartsWith string
}

func (f *RdsFilter) Match(value string) bool {
	return strings.HasPrefix(value, f.StartsWith)
}

type Rds struct {
	name    string
	tag     string
	session *rds.RDS
	filter  *RdsFilter
}

func NewRds(name, tag string, session *rds.RDS, clientFilter *RdsFilter) *Rds {
	return &Rds{
		name:    name,
		tag:     tag,
		session: session,
		filter:  clientFilter,
	}
}

func (r *Rds) Name() string {
	return r.name
}

func (r *Rds) Fetch() *[]Resource {
	instances, err := r.session.DescribeDBInstances(&rds.DescribeDBInstancesInput{})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println("fetcher.rds", aerr.Error())
			}
		} else {
			fmt.Println("fetcher.rds", err.Error())
		}
		return &[]Resource{}
	}

	resources := []Resource{}
	for _, i := range instances.DBInstances {
		if !r.filter.Match(*i.DBName) {
			continue
		}
		resources = append(resources, Resource{Name: i.DBName})
	}

	return &resources
}
