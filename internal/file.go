package internal

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/zkfmapf123/dobbyssm/utils"
	"gopkg.in/yaml.v3"
)

type AWSCredentials struct {
	Profile  string   `yaml:"profile"`
	Region   string   `yaml:"region"`
	Clusters []string `yaml:"clusters"`
}

const PATH = ".aws/ecsns.yaml"

func MustGetAWSCredentialsPath(path string) string {
	dir, err := os.UserHomeDir()
	if err != nil {
		utils.PanicRed(err)
	}

	return filepath.Join(dir, path)
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

func GetCredentialsFileUseParameter(suffixPath string) ([]string, error) {

	path := MustGetAWSCredentialsPath(suffixPath)

	if err := isExistsFile(path); err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var profiles []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			profiles = append(profiles, strings.Trim(line, "[]"))
		}
	}

	return profiles, nil
}

func GetCredentialFile() (AWSCredentials, error) {

	path := MustGetAWSCredentialsPath(PATH)

	if err := isExistsFile(path); err != nil {
		return AWSCredentials{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return AWSCredentials{}, err
	}

	defer file.Close()

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

func SetCredentialFile(profile string, region string, clusters []string) error {
	path := MustGetAWSCredentialsPath(PATH)

	creds := AWSCredentials{
		Profile:  profile,
		Region:   region,
		Clusters: clusters,
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
