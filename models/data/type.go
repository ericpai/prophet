package data

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type InvalidIaaSAccountError struct {
	Account  string `json:"account"`
	Service  string `json:"service"`
	Provider string `json:"provider"`
}

func (i InvalidIaaSAccountError) Error() string {
	return fmt.Sprintf(
		"invalid account config %s of %s.%s", i.Account, i.Provider, i.Service)
}

type InstancesOverview struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type InstanceOfferingView struct {
	OfferingTypes []ec2.OfferingTypeValues `json:"offering_types"`
	Offerings     []InstanceOffering       `json:"offerings"`
}

type InstanceOffering struct {
	Type   string `json:"type"`
	Counts []int  `json:"counts"`
}
