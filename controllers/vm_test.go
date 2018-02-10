package controllers

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/ericpai/prophet/libs"
	"github.com/ericpai/prophet/models"
	"github.com/ericpai/prophet/models/data"
)

type MockVMManager struct {
}

func (m *MockVMManager) OverviewInstances(account string) ([]data.InstancesOverview, error) {
	if account == "account1" {
		return []data.InstancesOverview{
			{
				Type:  "test.small",
				Count: 1,
			},
			{
				Type:  "test.medium",
				Count: 2,
			},
		}, nil
	} else {
		return nil, data.InvalidIaaSAccountError{
			Account:  account,
			Service:  "ec2",
			Provider: "test",
		}
	}
}

func TestGetVMInstancesHandler(t *testing.T) {
	secretary = Secretary{
		vmMamagers: map[string]models.VMManager{
			"test": &MockVMManager{},
		},
	}
	overviewSchema := map[string]interface{}{
		"type": "array",
		"items": map[string]interface{}{
			"type":                 "object",
			"additionalProperties": false,
			"properties": map[string]interface{}{
				"type": map[string]interface{}{
					"type": "string",
				},
				"count": map[string]interface{}{
					"type": "number",
				},
			},
			"required": []string{"type", "count"},
		},
	}
	testCases := []libs.APIAssertRequest{
		{
			Method: http.MethodGet,
			URL:    "/api/vm/overview",
			Values: url.Values{
				"account":  []string{"account1"},
				"provider": []string{"test"},
			},
			Schema: overviewSchema,
			Status: http.StatusOK,
		},
		{
			Method: http.MethodGet,
			URL:    "/api/vm/overview",
			Values: url.Values{
				"account":  []string{"account2"},
				"provider": []string{"test"},
			},
			Schema: overviewSchema,
			Status: http.StatusForbidden,
		},
		{
			Method: http.MethodGet,
			URL:    "/api/vm/overview",
			Values: url.Values{
				"account":  []string{"account1"},
				"provider": []string{"invalid_provider"},
			},
			Schema: overviewSchema,
			Status: http.StatusForbidden,
		},
		{
			Method: http.MethodGet,
			URL:    "/api/vm/overview",
			Values: url.Values{
				"account": []string{"account1"},
			},
			Schema: overviewSchema,
			Status: http.StatusUnprocessableEntity,
		},
		{
			Method: http.MethodGet,
			URL:    "/api/vm/overview",
			Values: url.Values{
				"provider": []string{"test"},
			},
			Schema: overviewSchema,
			Status: http.StatusUnprocessableEntity,
		},
	}

	for _, testCase := range testCases {
		testCase.AssertResponse(t, GetVMInstancesHandler)
	}

}
