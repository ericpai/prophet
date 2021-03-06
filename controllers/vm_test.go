package controllers

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/ericpai/prophet/libs"
	"github.com/ericpai/prophet/models"
	"github.com/ericpai/prophet/models/aws"
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
	}
	return nil, data.InvalidIaaSAccountError{
		Account:  account,
		Service:  "ec2",
		Provider: "test",
	}
}

func (m *MockVMManager) OverviewOfferings(account string) (data.InstanceOfferingView, error) {
	if account == "account1" {
		return data.InstanceOfferingView{
			OfferingTypes: []string{
				string(aws.OfferingTypeValuesOnDemand),
				string(ec2.OfferingTypeValuesAllUpfront),
				string(ec2.OfferingTypeValuesPartialUpfront),
				string(ec2.OfferingTypeValuesNoUpfront),
			},
			Offerings: []data.InstanceOffering{
				{
					Type:   "m1.xlarge",
					Counts: []int{1, 2, 3, 4},
				},
				{
					Type:   "m2.xlarge",
					Counts: []int{4, 3, 2, 1},
				},
			},
		}, nil
	}
	return data.InstanceOfferingView{}, data.InvalidIaaSAccountError{
		Account:  account,
		Service:  "ec2",
		Provider: "test",
	}
}

func (m *MockVMManager) OverviewStorage(account string) (data.VMStorage, error) {
	if account == "account1" {
		return data.VMStorage{
			Currency: "￥",
			Unit:     "GB",
			Volumes: map[string]data.VMStorageVolume{
				"g2": {
					Type:   "g2",
					Cost:   100.0,
					Amount: 250,
				},
			},
		}, nil
	}
	return data.VMStorage{}, data.InvalidIaaSAccountError{
		Account:  account,
		Service:  "ec2",
		Provider: "test",
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

func TestGetVMOfferingsHandler(t *testing.T) {
	secretary = Secretary{
		vmMamagers: map[string]models.VMManager{
			"test": &MockVMManager{},
		},
	}
	offeringsSchema := map[string]interface{}{
		"type":                 "object",
		"additionalProperties": false,
		"properties": map[string]interface{}{
			"offering_types": map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "string",
				},
			},
			"offerings": map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type":                 "object",
					"additionalProperties": false,
					"properties": map[string]interface{}{
						"type": map[string]interface{}{
							"type": "string",
						},
						"counts": map[string]interface{}{
							"type": "array",
							"items": map[string]interface{}{
								"type": "number",
							},
						},
					},
					"required": []string{"type", "counts"},
				},
			},
		},
		"required": []string{"offering_types", "offerings"},
	}
	testCases := []libs.APIAssertRequest{
		{
			Method: http.MethodGet,
			URL:    "/api/vm/offerings",
			Values: url.Values{
				"account":  []string{"account1"},
				"provider": []string{"test"},
			},
			Schema: offeringsSchema,
			Status: http.StatusOK,
		},
		{
			Method: http.MethodGet,
			URL:    "/api/vm/offerings",
			Values: url.Values{
				"account":  []string{"account2"},
				"provider": []string{"test"},
			},
			Schema: offeringsSchema,
			Status: http.StatusForbidden,
		},
		{
			Method: http.MethodGet,
			URL:    "/api/vm/offerings",
			Values: url.Values{
				"account":  []string{"account1"},
				"provider": []string{"invalid_provider"},
			},
			Schema: offeringsSchema,
			Status: http.StatusForbidden,
		},
		{
			Method: http.MethodGet,
			URL:    "/api/vm/offerings",
			Values: url.Values{
				"account": []string{"account1"},
			},
			Schema: offeringsSchema,
			Status: http.StatusUnprocessableEntity,
		},
		{
			Method: http.MethodGet,
			URL:    "/api/vm/offerings",
			Values: url.Values{
				"provider": []string{"test"},
			},
			Schema: offeringsSchema,
			Status: http.StatusUnprocessableEntity,
		},
	}

	for _, testCase := range testCases {
		testCase.AssertResponse(t, GetVMOfferingsHandler)
	}

}

func TestGetVMStorageHandler(t *testing.T) {
	secretary = Secretary{
		vmMamagers: map[string]models.VMManager{
			"test": &MockVMManager{},
		},
	}
	storageSchema := map[string]interface{}{
		"type":                 "object",
		"additionalProperties": false,
		"properties": map[string]interface{}{
			"currency": map[string]interface{}{
				"type": "string",
			},
			"unit": map[string]interface{}{
				"type": "string",
			},
			"volumes": map[string]interface{}{
				"type": "object",
				"patternProperties": map[string]interface{}{
					"^.*$": map[string]interface{}{
						"type":                 "object",
						"additionalProperties": false,
						"properties": map[string]interface{}{
							"type": map[string]interface{}{
								"type": "string",
							},
							"amount": map[string]interface{}{
								"type": "number",
							},
							"cost": map[string]interface{}{
								"type": "number",
							},
						},
						"required": []string{"amount", "cost", "type"},
					},
				},
			},
		},
		"required": []string{"currency", "unit", "volumes"},
	}
	testCases := []libs.APIAssertRequest{
		{
			Method: http.MethodGet,
			URL:    "/api/vm/storage",
			Values: url.Values{
				"account":  []string{"account1"},
				"provider": []string{"test"},
			},
			Schema: storageSchema,
			Status: http.StatusOK,
		},
		// {
		// 	Method: http.MethodGet,
		// 	URL:    "/api/vm/storage",
		// 	Values: url.Values{
		// 		"account":  []string{"account2"},
		// 		"provider": []string{"test"},
		// 	},
		// 	Schema: storageSchema,
		// 	Status: http.StatusForbidden,
		// },
		// {
		// 	Method: http.MethodGet,
		// 	URL:    "/api/vm/storage",
		// 	Values: url.Values{
		// 		"account":  []string{"account1"},
		// 		"provider": []string{"invalid_provider"},
		// 	},
		// 	Schema: storageSchema,
		// 	Status: http.StatusForbidden,
		// },
		// {
		// 	Method: http.MethodGet,
		// 	URL:    "/api/vm/storage",
		// 	Values: url.Values{
		// 		"account": []string{"account1"},
		// 	},
		// 	Schema: storageSchema,
		// 	Status: http.StatusUnprocessableEntity,
		// },
		// {
		// 	Method: http.MethodGet,
		// 	URL:    "/api/vm/storage",
		// 	Values: url.Values{
		// 		"provider": []string{"test"},
		// 	},
		// 	Schema: storageSchema,
		// 	Status: http.StatusUnprocessableEntity,
		// },
	}

	for _, testCase := range testCases {
		testCase.AssertResponse(t, GetVMStorageHandler)
	}

}
