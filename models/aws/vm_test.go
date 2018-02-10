package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/ec2iface"
	"github.com/ericpai/prophet/models/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockEC2Client struct {
	ec2iface.EC2API
}

func (m mockEC2Client) DescribeInstancesPagesWithContext(
	ctx aws.Context,
	input *ec2.DescribeInstancesInput,
	fn func(*ec2.DescribeInstancesOutput, bool) bool,
	opts ...aws.Option) error {
	output := &ec2.DescribeInstancesOutput{
		Reservations: []ec2.RunInstancesOutput{
			{
				Instances: []ec2.Instance{
					{
						InstanceType: ec2.InstanceTypeC1Xlarge,
					},
					{
						InstanceType: ec2.InstanceTypeC32xlarge,
					},
				},
			},
			{
				Instances: []ec2.Instance{
					{
						InstanceType: ec2.InstanceTypeC1Medium,
					},
					{
						InstanceType: ec2.InstanceTypeC32xlarge,
					},
				},
			},
		},
	}
	fn(output, true)
	return nil
}

func TestOverviewInstances(t *testing.T) {
	expected := []data.InstancesOverview{
		{
			Type:  "c1.medium",
			Count: 1,
		},
		{
			Type:  "c1.xlarge",
			Count: 1,
		},
		{
			Type:  "c3.2xlarge",
			Count: 2,
		},
	}
	m := &VMManager{
		api: map[string]ec2iface.EC2API{
			"mock": &mockEC2Client{},
		},
	}
	actual, err := m.OverviewInstances("mock")
	assert.NoError(t, err)
	assert.EqualValues(t, expected, actual)
	actual, err = m.OverviewInstances("notfound")
	assert.Error(t, err)
	assert.Nil(t, actual)
}
