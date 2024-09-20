package internal

import (
	"testing"
)

// func afterHooks() {
// 	fmt.Println("Hooks Clear")
// 	err := os.Remove(MustGetAWSCredentialsPath(PATH))
// 	if err != nil {
// 		utils.PanicRed(err)
// 	}
// }

func Test_GetCredentialFile(t *testing.T) {

	// _, notExistError := GetCredentialFile()
	// assert.Equal(t, strings.Contains(notExistError.Error(), "no such file or directory"), true)

	// proifle, region, cluster := "dobby", "ap-northeast-2", []string{"dobby-cluster-1", "dobby-cluster-2"}
	// err := SetCredentialFile(proifle, region, cluster)
	// assert.Equal(t, err, nil)

	// creds, _ := GetCredentialFile()

	// assert.Equal(t, creds.Profile, proifle)
	// assert.Equal(t, creds.Region, region)
	// assert.Equal(t, creds.Clusters, cluster)

	// afterHooks()
}
