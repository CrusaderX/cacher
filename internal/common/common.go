package common

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/rds"
)

type Filter struct {
	Key		string `json:"Key"`
	Value string `json:"Value"`
}

type Ec2 struct {
	Filters []Filter `json:"Filters"`
	Method
}

func (c *EC2) Describe(input *DescribeInstancesInput) (*DescribeDBInstancesOutput, error) {}

type Rds struct {
	Filters []Filter `json:"Filters"`
	Method
}

func (c *RDS) Describe(input *DescribeDBInstancesInput) (*DescribeDBInstancesOutput, error) {}
