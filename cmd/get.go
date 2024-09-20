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

			t := utils.NewTerminal("").Clear()
			awsCreds, err := internal.GetCredentialFile()
			if err != nil {
				utils.PanicRed(err)
			}

			awsConn := internal.NewAWS(awsCreds.Profile, awsCreds.Region)
			awsConn.SetClusters(awsCreds.Clusters)

			if len(args) == 0 {
				shortHandTableWriter(t)
				utils.WarningYellow(errors.New("Please enter the resource you want to get"))
			}

			apiResources, command := base.GetAPIResources(), args[0]

			// clusters
			if utils.IncludeString(apiResources.Cluster, command) {
				utils.MustCheckError(getCluster(awsConn, t))
				return
			}

			// services
			if utils.IncludeString(apiResources.Services, command) {
				utils.MustCheckError(getService(awsConn, t))
				return
			}

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
	headers, values, err := awsConn.GetECSClusterDetails()
	if err != nil {
		utils.PanicRed(err)
	}

	t.TableWriter(headers, values)
	return nil
}

func getService(awsConn internal.AWSParams, t utils.Termianl) error {

	headers, values, err := awsConn.GetECSService()
	if err != nil {
		utils.PanicRed(err)
	}

	t.TableWriter(headers, values)

	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)
}
