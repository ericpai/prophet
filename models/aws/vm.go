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

type VMManager struct {
	api map[string]ec2iface.EC2API
}

func NewVMManager() *VMManager {
	apis, _ := iaas.InitAWSEC2Client()
	return &VMManager{
		api: apis,
	}
}

func (m *VMManager) OverviewInstances(account string) ([]data.InstancesOverview, error) {
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
