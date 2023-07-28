package main

import (
	"encoding/json"
	"fmt"

	"github.com/CrusaderX/cacher/internal/fetcher"
	"github.com/CrusaderX/cacher/internal/registry"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	// "github.com/aws/aws-sdk-go/service/rds"
)

func main() {
	reg := registry.NewFetcherRegistry()
	defer reg.Close()

	awssession, _ := session.NewSession()
	ec2session := ec2.New(awssession)
	// rdssession := rds.New(awssession)

	reg.Register(fetcher.NewEc2("EC2", "enabled", ec2session))
	// reg.Register(fetcher.NewRds("RDS", "enabled", rdssession, &fetcher.RdsFilter{StartsWith: "dev"}))

	go reg.Fetch()

	for r := range reg.Results() {
		response, err := json.Marshal(r)

		if err != nil {
			return
		}
		fmt.Println(string(response))
	}
}
