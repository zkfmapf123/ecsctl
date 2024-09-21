package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/dobbyssm/internal"
	"github.com/zkfmapf123/dobbyssm/utils"
	gj "github.com/zkfmapf123/go-js-utils"
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

type ecsExecItem struct {
	taskId        string
	containerName string
}

type ecsAttr map[string][]ecsExecItem

const (
	CLUSTER_NAME_INDEX   = iota // 0
	TASK_ID_INDEX               // 1
	CONTAINER_NAME_INDEX        // 2

	EXEC_INDEX = 8
)

var (
	execCmd = &cobra.Command{
		Use:   "exec",
		Short: "exec ecs container",
		Long:  "exec ecs container",
		Run: func(cmd *cobra.Command, args []string) {
			awsCreds, err := internal.GetCredentialFile()
			if err != nil {
				utils.PanicRed(err)
			}

			awsConn := internal.NewAWS(awsCreds.Profile, awsCreds.Region)
			awsConn.SetClusters(awsCreds.Clusters)

			t := utils.NewTerminal("")
			t.Clear()

			// 1. select task
			_, ecsContainers, err := awsConn.GetECSContainers()
			if err != nil {
				utils.PanicRed(err)
			}

			ecsAtrrParams := ecsAttr{}
			for _, ecsAttr := range ecsContainers {

				// must be enable exec command
				if ecsAttr[EXEC_INDEX] != "true" {
					continue
				}

				clusterName := ecsAttr[CLUSTER_NAME_INDEX]
				taskId := ecsAttr[TASK_ID_INDEX]
				containerName := ecsAttr[CONTAINER_NAME_INDEX]

				item := ecsExecItem{
					taskId:        taskId,
					containerName: containerName,
				}

				if v, isOk := ecsAtrrParams[clusterName]; isOk {
					// value exists
					ecsAtrrParams[clusterName] = append(v, item)

				} else {
					// not exists
					ecsAtrrParams[clusterName] = []ecsExecItem{item}
				}
			}

			clusterNames := gj.OKeys(ecsAtrrParams)
			selectClusterName := t.SelectOne("Select ECS Cluster", clusterNames)

			taskIdMap := map[string][]string{}
			for _, ecsAttr := range ecsAtrrParams[selectClusterName] {
				if v, isOk := taskIdMap[ecsAttr.taskId]; isOk {
					taskIdMap[ecsAttr.taskId] = append(v, ecsAttr.containerName)
				} else {
					taskIdMap[ecsAttr.taskId] = []string{ecsAttr.containerName}
				}
			}

			// 2. select TaskId
			taskIds := gj.OKeys(taskIdMap)
			selectTaskId := t.SelectOne("Select ECS Task", taskIds)

			// 3. select container name
			selectContainerName := t.SelectOne("Select Container Name", taskIdMap[selectTaskId])

			params := execParameter{
				clusterName:   selectClusterName,
				taskId:        selectTaskId,
				containerName: selectContainerName,
				region:        awsCreds.Region,
				profile:       awsCreds.Profile,
			}

			connectSSM(params)
		},
	}
)

func connectSSM(params execParameter) {
	// ssm connect (call process)
	call := exec.Command(
		"aws",
		"ecs",
		"execute-command",
		"--profile",
		params.profile,
		"--region",
		params.region,
		"--cluster",
		params.clusterName,
		"--task",
		params.taskId,
		"--container",
		params.containerName,
		"--command",
		"/bin/sh",
		"--interactive",
	)
	call.Stderr = os.Stderr
	call.Stdout = os.Stdout
	call.Stdin = os.Stdin
	call.Run()
}

func init() {

	// 2024.9.21 선택으로 수정 (cluster, taskId, container)
	// execCmd.Flags().StringP("cluster", "c", "", "[Required] cluster name")
	// execCmd.Flags().StringP("task", "t", "", "[Required] task id")
	// execCmd.Flags().StringP("container", "o", "", "[Required] container name")

	// viper.BindPFlag("cluster", execCmd.Flags().Lookup("cluster"))
	// viper.BindPFlag("task", execCmd.Flags().Lookup("task"))
	// viper.BindPFlag("container", execCmd.Flags().Lookup("container"))

	rootCmd.AddCommand(execCmd)
}
