package main

import (
	"fmt"

	"github.com/CrusaderX/cacher/internal/fetcher"
	"github.com/CrusaderX/cacher/internal/registry"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	reg := registry.NewFetcherRegistry()
	defer reg.Close()

	ec2session := ec2.New(session.New())

	reg.Register(fetcher.NewEc2("EC2", "enabled", ec2session))
	reg.Register(fetcher.NewRds("RDS", "enabled"))

	go reg.Fetch()

	for r := range reg.Results() {
		fmt.Println(r)
	}
}
