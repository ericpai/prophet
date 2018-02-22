package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/ec2iface"
	"github.com/ericpai/prophet/iaas"
	"github.com/ericpai/prophet/models/data"
	"sort"
	"time"
)

const OfferingTypeValuesOnDemand ec2.OfferingTypeValues = "On Demand"

type VMManager struct {
	api map[string]ec2iface.EC2API
}

func NewVMManager() *VMManager {
	apis, _ := iaas.InitAWSEC2Client()
	return &VMManager{
		api: apis,
	}
}

func (m *VMManager) OverviewInstances(account string) (
	[]data.InstancesOverview, error) {
	api, exist := m.api[account]
	if !exist {
		return nil, data.InvalidIaaSAccountError{
			Account:  account,
			Service:  "ec2",
			Provider: "aws",
		}
	}
	var overview []data.InstancesOverview
	typeMap := make(map[string]int)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err := api.DescribeInstancesPagesWithContext(
		ctx,
		&ec2.DescribeInstancesInput{},
		func(output *ec2.DescribeInstancesOutput, lastPage bool) bool {
			for _, out := range output.Reservations {
				for _, inst := range out.Instances {
					typeMap[string(inst.InstanceType)]++
				}
			}
			return true
		},
	)
	if err != nil {
		return nil, err
	}
	for k, v := range typeMap {
		overview = append(overview, data.InstancesOverview{
			Type:  k,
			Count: v,
		})
	}
	sort.Slice(overview, func(i, j int) bool {
		return overview[i].Type < overview[j].Type
	})
	return overview, nil
}

func (m *VMManager) OverviewOfferings(account string) (
	data.InstanceOfferingView, error) {
	rv := data.InstanceOfferingView{
		OfferingTypes: []ec2.OfferingTypeValues{
			OfferingTypeValuesOnDemand,
			ec2.OfferingTypeValuesHeavyUtilization,
			ec2.OfferingTypeValuesLightUtilization,
			ec2.OfferingTypeValuesMediumUtilization,
			ec2.OfferingTypeValuesAllUpfront,
			ec2.OfferingTypeValuesNoUpfront,
		},
		Offerings: []data.InstanceOffering{},
	}
	api, exist := m.api[account]
	if !exist {
		return rv, data.InvalidIaaSAccountError{
			Account:  account,
			Service:  "ec2",
			Provider: "aws",
		}
	}

	stateFilterName := "state"
	req := api.DescribeReservedInstancesRequest(
		&ec2.DescribeReservedInstancesInput{
			Filters: []ec2.Filter{
				{
					Name: &stateFilterName,
					Values: []string{
						"active",
					},
				},
			},
		},
	)
	output, err := req.Send()
	if err != nil {
		return rv, err
	}
	instances, err := m.OverviewInstances(account)
	if err != nil {
		return rv, err
	}
	tmpMap := make(map[string]map[ec2.OfferingTypeValues]int)
	for _, inst := range instances {
		tmpMap[inst.Type] = map[ec2.OfferingTypeValues]int{
			OfferingTypeValuesOnDemand:              inst.Count,
			ec2.OfferingTypeValuesHeavyUtilization:  0,
			ec2.OfferingTypeValuesLightUtilization:  0,
			ec2.OfferingTypeValuesMediumUtilization: 0,
			ec2.OfferingTypeValuesAllUpfront:        0,
			ec2.OfferingTypeValuesNoUpfront:         0,
		}
	}
	for _, ri := range output.ReservedInstances {
		instTypeStr := (string)(ri.InstanceType)
		if _, exist := tmpMap[instTypeStr]; !exist {
			tmpMap[(string)(ri.InstanceType)] = map[ec2.OfferingTypeValues]int{
				OfferingTypeValuesOnDemand:              0,
				ec2.OfferingTypeValuesHeavyUtilization:  0,
				ec2.OfferingTypeValuesLightUtilization:  0,
				ec2.OfferingTypeValuesMediumUtilization: 0,
				ec2.OfferingTypeValuesAllUpfront:        0,
				ec2.OfferingTypeValuesNoUpfront:         0,
			}
		}
		tmpMap[instTypeStr][OfferingTypeValuesOnDemand] -=
			(int)(*ri.InstanceCount)
		tmpMap[instTypeStr][ri.OfferingType] += (int)(*ri.InstanceCount)
	}

	for k, v := range tmpMap {
		countSlice := make([]int, len(rv.OfferingTypes))
		for i, o := range rv.OfferingTypes {
			countSlice[i] = v[o]
		}
		rv.Offerings = append(rv.Offerings, data.InstanceOffering{
			Type:   k,
			Counts: countSlice,
		})
	}
	return rv, nil
}
