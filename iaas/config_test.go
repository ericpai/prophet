package iaas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	expected := ProphetConfig{
		Currency: "CNY",
		AWS: ProphetAWSConfig{
			Accounts: []AWSAccount{
				{
					Name:            "test_a",
					AccessKeyID:     "key_test_a",
					SecretAccessKey: "secret_key_a",
					Region:          "cn-north-1",
					Services:        []string{"ec2", "s3"},
				},
			},
		},
	}
	result, err := loadConfig("conf/example-account-secret.yaml")
	assert.NoError(t, err)
	assert.Equal(t, expected, *result)

}

func TestHasServicePrivilege(t *testing.T) {
	a := AWSAccount{
		Services: []string{"ec2", "s3"},
	}
	assert.True(t, a.hasServicePrivilege("ec2"))
	assert.True(t, a.hasServicePrivilege("s3"))
	assert.False(t, a.hasServicePrivilege("sqs"))
}
