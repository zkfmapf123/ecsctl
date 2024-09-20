package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zkfmapf123/dobbyssm/utils"
)

var (
	apiResourceCmd = &cobra.Command{
		Use:   "api-resources",
		Short: "Description ecsctl api-resource",
		Long:  "Description ecsctl api-resource",
		Run: func(cmd *cobra.Command, args []string) {

			t := utils.NewTerminal("")

			values := [][]string{
				{"api-resources", "api-resources", "Description Resource"},
				{"clusters", "cl", "ECS Cluster"},
				{"services", "s", "ECS Service"},
				{"containers", "c", "Container in ECS Service"},
				{"tasks", "t", "ECS Task"},
				{"alb", "al", "ECS ALB Information"},
			}

			t.TableWriter([]string{"Name", "SHORTHAND", "Description"}, values)
		},
	}
)

func init() {
	rootCmd.AddCommand(apiResourceCmd)
}
