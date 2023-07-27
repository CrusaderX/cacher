package main

import (
	"encoding/json"
	"fmt"
	"github.com/CrusaderX/cacher"
)

func main() {
	filters := []Filter{
		{Key: "namespace", Value: "default"},
		{Key: "scheduler", Value: "enabled"},
	}
	
	ec2 := Ec2{
		Filters: filters,
	}
	rds := Rds{
		Filters: filters,
	}

	c1 := make(chan Ec2)
	c2 := make(chan Rds)

	go func() { sleep 300; c1 <- ec2 }
	
}
