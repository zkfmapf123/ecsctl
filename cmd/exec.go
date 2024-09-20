package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zkfmapf123/dobbyssm/internal"
	"github.com/zkfmapf123/dobbyssm/utils"
)

////////////////////////////////////////////////////////////////////////
// aws ecs execute-command --profile PROFLE_NAME --region REGION_NAME\
// --cluster CLUSTER_NAME \
// --task TASK_ID \
// --container CONTAINER_NAME \
// --command "/bin/sh" \
// --interactive
////////////////////////////////////////////////////////////////////////

type execParameter struct {
	clusterName   string
	taskId        string
	containerName string
	region        string
	profile       string
}

var (
	execCmd = &cobra.Command{
		Use:   "exec",
		Short: "exec ecs container",
		Long:  "exec ecs container",
		Run: func(cmd *cobra.Command, args []string) {
			// get Credential
			awsCreds, err := internal.GetCredentialFile()
			if err != nil {
				utils.PanicRed(err)
			}

			// parameter
			execParams := inspectParameters()
			execParams.region = awsCreds.Region
			execParams.profile = awsCreds.Profile

			// ssm connect (call process)

		},
	}
)

func inspectParameters() execParameter {

	cluster, task, container := viper.GetString(
		"cluster",
	), viper.GetString(
		"task",
	), viper.GetString(
		"container",
	)

	if cluster == "" {
		utils.PanicRed(errors.New("cluster is required"))
	}

	if task == "" {
		utils.PanicRed(errors.New("task is required"))
	}

	if container == "" {
		utils.PanicRed(errors.New("container is required"))
	}

	return execParameter{
		clusterName:   cluster,
		taskId:        task,
		containerName: container,
	}

}

func init() {

	execCmd.Flags().StringP("cluster", "c", "", "[Required] cluster name")
	execCmd.Flags().StringP("task", "t", "", "[Required] task id")
	execCmd.Flags().StringP("container", "o", "", "[Required] container name")

	viper.BindPFlag("cluster", execCmd.Flags().Lookup("cluster"))
	viper.BindPFlag("task", execCmd.Flags().Lookup("task"))
	viper.BindPFlag("container", execCmd.Flags().Lookup("container"))

	rootCmd.AddCommand(execCmd)
}
