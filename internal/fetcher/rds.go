package fetcher

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rds"
)

type Rds struct {
	name    string
	tag     string
	session *rds.RDS
}

func NewRds(name, tag string, session *rds.RDS) *Rds {
	return &Rds{
		name:    name,
		tag:     tag,
		session: session,
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

	namespaces := make(map[string][]string)
	for _, i := range instances.DBInstances {
		var namespace *string
		isSchedulerEnabled := false
		for _, t := range i.TagList {
			if *t.Key == "scheduler" && *t.Value == "enabled" {
				isSchedulerEnabled = true
				continue
			}
			if *t.Key == "namespace" {
				namespace = t.Value
			}
		}
		if !isSchedulerEnabled {
			continue
		}
		if namespace == nil {
			fmt.Printf("no namespace for rds %s. skipping.\n", *i.DBInstanceIdentifier)
			continue
		}
		namespaces[*namespace] = append(namespaces[*namespace], *i.DBName)
	}

	var resources []Resource
	for namespace, instanceIds := range namespaces {
		resources = append(resources, Resource{
			Namespace: map[string][]string{
				namespace: instanceIds,
			},
		})
	}

	return &resources
}
