package internal

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkfmapf123/dobbyssm/utils"
)

func afterHooks() {
	fmt.Println("Hooks Clear")
	err := os.Remove(mustGetAWSCredentialsPath())
	if err != nil {
		utils.PanicRed(err)
	}
}

func Test_GetCredentialFile(t *testing.T) {

	_, notExistError := GetCredentialFile()
	assert.Equal(t, strings.Contains(notExistError.Error(), "no such file or directory"), true)

	proifle, region, cluster := "dobby", "ap-northeast-2", "dobby-cluster"
	err := SetCredentialFile(proifle, region, cluster)
	assert.Equal(t, err, nil)

	creds, _ := GetCredentialFile()
	fmt.Println(creds)

	assert.Equal(t, creds.Profile, proifle)
	assert.Equal(t, creds.Region, region)
	assert.Equal(t, creds.Cluster, cluster)

	afterHooks()
}
