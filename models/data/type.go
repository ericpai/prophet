package data

import (
	"fmt"
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
	OfferingTypes []string           `json:"offering_types"`
	Offerings     []InstanceOffering `json:"offerings"`
}

type InstanceOffering struct {
	Type   string `json:"type"`
	Counts []int  `json:"counts"`
}

type VMStorage struct {
	Unit     string                     `json:"unit"`
	Currency string                     `json:"currency"`
	Volumes  map[string]VMStorageVolume `json:"volumes"`
}

type VMStorageVolume struct {
	Type   string  `json:"type"`
	Cost   float64 `json:"cost"`
	Amount int     `json:"amount"`
}
