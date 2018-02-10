package models

import (
	"github.com/ericpai/prophet/models/aws"
	"github.com/ericpai/prophet/models/data"
)

type VMManager interface {
	OverviewInstances(account string) ([]data.InstancesOverview, error)
}

func GetVMManagers() map[string]VMManager {
	return map[string]VMManager{
		"aws": aws.NewVMManager(),
	}
}
