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
	return fmt.Sprintf("invalid account config %s of %s.%s", i.Account, i.Provider, i.Service)
}

type InstancesOverview struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}
