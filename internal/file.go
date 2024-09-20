package internal

import (
	"io"
	"os"
	"path/filepath"

	"github.com/zkfmapf123/dobbyssm/utils"
	"gopkg.in/yaml.v3"
)

type AWSCredentials struct {
	Profile string `yaml:"profile"`
	Region  string `yaml:"region"`
	Cluster string `yaml:"cluster"`
}

const PATH = ".aws/ecsns.yaml"

func mustGetAWSCredentialsPath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		utils.PanicRed(err)
	}

	return filepath.Join(dir, PATH)
}

func isExistsFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err == nil {
		return err
	}

	if os.IsNotExist(err) {
		return nil
	}

	return nil
}

func GetCredentialFile() (AWSCredentials, error) {

	path := mustGetAWSCredentialsPath()

	if err := isExistsFile(path); err != nil {
		return AWSCredentials{}, err
	}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return AWSCredentials{}, err
	}

	fb, err := io.ReadAll(file)
	if err != nil {
		return AWSCredentials{}, err
	}

	var creds AWSCredentials
	if err := yaml.Unmarshal(fb, &creds); err != nil {
		return AWSCredentials{}, err
	}

	return creds, nil
}

func SetCredentialFile(profile string, region string, cluster string) error {
	path := mustGetAWSCredentialsPath()

	creds := AWSCredentials{
		Profile: profile,
		Region:  region,
		Cluster: cluster,
	}

	yamlBytes, err := yaml.Marshal(creds)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, yamlBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
