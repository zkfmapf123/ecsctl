package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/dobbyssm/internal"
	"github.com/zkfmapf123/dobbyssm/utils"
)

var RegionList = []string{
	"af-south-1",
	"ap-east-1",
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-northeast-3",
	"ap-south-1",
	"ap-south-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-southeast-3",
	"ap-southeast-4",
	"ap-southeast-5",
	"ca-central-1",
	"ca-west-1",
	"cn-north-1",
	"cn-northwest-1",
	"eu-central-1",
	"eu-central-2",
	"eu-north-1",
	"eu-south-1",
	"eu-south-2",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	"il-central-1",
	"me-central-1",
	"me-south-1",
	"sa-east-1",
	"us-east-1",
	"us-east-2",
	"us-gov-east-1",
	"us-gov-west-1",
	"us-west-1",
	"us-west-2",
}

var (
	ecsnsCmd = &cobra.Command{
		Use:   "ecsns",
		Short: "[Required] setting aws credentials",
		Long:  "[Required] setting aws credentials",
		Run: func(cmd *cobra.Command, args []string) {
			t := utils.NewTerminal(internal.MustGetAWSCredentialsPath("")).Clear()

			profiles, err := internal.GetCredentialsFileUseParameter(".aws/credentials")
			fmt.Println(profiles, err)

			creds := internal.AWSCredentials{}

			// 1. get aws profile
			creds.Profile = t.SelectOne("Select AWS Profile", profiles)
			t.Clear()

			// 2. get region
			creds.Region = t.SelectOne("Select AWS Region", RegionList)
			t.Clear()

			awsConn := internal.NewAWS(creds.Profile, creds.Region)
			clusters, err := awsConn.GetECSCluster()
			if err != nil {
				utils.PanicRed(err)
			}
			t.Clear()

			// 3. get cluster
			creds.Clusters = t.SelectMultiple("Select ECS Cluster", clusters)
			t.Clear()

			// 4. save credentials
			err = internal.SetCredentialFile(creds.Profile, creds.Region, creds.Clusters)
			if err != nil {
				utils.PanicRed(err)
			}
			t.Clear()

			utils.SuccessGreen("Success to save credentials")
		},
	}
)

func init() {

	rootCmd.AddCommand(ecsnsCmd)
}
