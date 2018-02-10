package controllers

import (
	"fmt"

	"github.com/ericpai/prophet/models"
)

type UnsupportedIaaSProviderError struct {
	iaasProvider string
}

func (e UnsupportedIaaSProviderError) Error() string {
	return fmt.Sprintf("unsupported IaaS provider: %s", e.iaasProvider)
}

type Secretary struct {
	vmMamagers map[string]models.VMManager
}

var secretary Secretary

func InitSecretary() {
	secretary = Secretary{
		vmMamagers: models.GetVMManagers(),
	}
}

func getVMManager(iaasProvider string) (models.VMManager, error) {
	manager, exist := secretary.vmMamagers[iaasProvider]
	if !exist {
		return nil, UnsupportedIaaSProviderError{iaasProvider: iaasProvider}
	}
	return manager, nil
}
