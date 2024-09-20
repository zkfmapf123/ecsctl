package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/dobbyssm/base"
	"github.com/zkfmapf123/dobbyssm/internal"
	"github.com/zkfmapf123/dobbyssm/utils"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get ecsctl resources",
		Long:  "get ecsctl resources",
		Run: func(cmd *cobra.Command, args []string) {

			t := utils.NewTerminal("")
			awsCreds, err := internal.GetCredentialFile()
			if err != nil {
				utils.PanicRed(err)
			}

			awsConn := internal.NewAWS(awsCreds.Profile, awsCreds.Region)

			if len(args) == 0 {
				shortHandTableWriter(t)
				utils.WarningYellow(errors.New("Please enter the resource you want to get"))
			}

			apiResources, command := base.GetAPIResources(), args[0]

			// clusters
			if utils.IncludeString(apiResources.Cluster, command) {
				utils.MustCheckError(getCluster(awsConn, t))
			}

			// // services
			// if utils.IncludeString(apiResources.Services, command) {

			// }

			// // containers
			// if utils.IncludeString(apiResources.Containers, command) {

			// }

			// // tasks
			// if utils.IncludeString(apiResources.Tasks, command) {

			// }

			// // alb
			// if utils.IncludeString(apiResources.Alb, command) {

			// }

			shortHandTableWriter(t)
			utils.WarningYellow(errors.New(""))
		},
	}
)

func getCluster(awsConn internal.AWSParams, t utils.Termianl) error {
	clusters, err := awsConn.GetECSCluster()
	if err != nil {
		utils.PanicRed(err)
	}

	clusterValues := make([][]string, len(clusters))
	for i, v := range clusters {
		clusterValues[i] = []string{v}
	}

	t.TableWriter([]string{"Cluster Name"}, clusterValues)
	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)
}
