package iaas

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/ec2iface"
)

func InitAWSEC2Client() (map[string]ec2iface.EC2API, error) {
	ec2Clients := make(map[string]ec2iface.EC2API)
	for _, account := range prophetConfig.AWS.Accounts {
		cfg, err := external.LoadDefaultAWSConfig()
		if err != nil {
			return nil, err
		}
		if account.hasServicePrivilege("ec2") {
			cfg.Credentials = aws.NewStaticCredentialsProvider(account.AccessKeyID, account.SecretAccessKey, "")
			cfg.Region = account.Region
			ec2Clients[account.Name] = ec2.New(cfg)
		}
	}
	return ec2Clients, nil
}
