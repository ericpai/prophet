package iaas

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ProphetConfig struct {
	Currency string           `yaml:"currency"`
	AWS      ProphetAWSConfig `yaml:"aws"`
}

type ProphetAWSConfig struct {
	Accounts []AWSAccount
}

type AWSAccount struct {
	Name            string   `yaml:"name"`
	AccessKeyID     string   `yaml:"access_key_id"`
	SecretAccessKey string   `yaml:"secret_access_key"`
	Region          string   `yaml:"region"`
	Services        []string `yaml:"services"`
}

func (a *AWSAccount) hasServicePrivilege(service string) bool {
	for _, v := range a.Services {
		if v == service {
			return true
		}
	}
	return false
}

var prophetConfig *ProphetConfig

func loadConfig(filename string) (*ProphetConfig, error) {
	pc := &ProphetConfig{}
	yamlContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(yamlContent, pc); err != nil {
		return nil, err
	}
	return pc, nil
}

func InitConfig(filename string) {
	var err error
	prophetConfig, err = loadConfig(filename)
	if err != nil {
		panic(err)
	}
}
