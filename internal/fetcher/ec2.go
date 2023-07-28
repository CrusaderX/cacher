package fetcher

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Ec2 struct {
	name    string
	tag     string
	session *ec2.EC2
}

func NewEc2(name, tag string, session *ec2.EC2) *Ec2 {
	return &Ec2{
		name:    name,
		tag:     tag,
		session: session,
	}
}

func (e *Ec2) Name() string {
	return e.name
}

func (e *Ec2) Fetch() *[]Resource {
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:scheduler"),
				Values: []*string{
					aws.String(e.tag),
				},
			},
		},
	}
	r, err := e.session.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println("fetcher.ec2", aerr.Error())
			}
		} else {
			fmt.Println("fetcher.ec2", err.Error())
		}
		return &[]Resource{}
	}

	resources := []Resource{}
	for _, res := range r.Reservations {
		for _, instance := range res.Instances {
			var instanceName *string
			for _, t := range instance.Tags {
				if *t.Key == "Name" {
					instanceName = t.Value
				}
			}
			if instanceName == nil {
				fmt.Println("fetcher.ec2", "cannot find instance name", instance)
				continue
			}
			resources = append(resources, Resource{Tags: &Tags{Name: *instanceName}})
		}
	}

	return &resources
}
