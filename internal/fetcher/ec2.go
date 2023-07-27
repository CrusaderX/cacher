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

func (e *Ec2) Fetch() []string {
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
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}
	fmt.Println(r)
	return []string{"3", "4"}
}
