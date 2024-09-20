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

			t := utils.NewTerminal("").Clear()
			shortHandTableWriter(t)
		},
	}
)

func shortHandTableWriter(t utils.Termianl) {

	values := [][]string{
		{"api-resources", "api-resources", "Description Resource"},
		{"clusters", "c, clu", "ECS Cluster"},
		{"pods", "s, svc", "ECS Service"},
		{"containers", "co, con", "Container in ECS Service"},
		{"alb", "al, alb", "ECS ALB Information"},
	}

	t.TableWriter([]string{"Name", "SHORTHAND", "Description"}, values)
}

func init() {
	rootCmd.AddCommand(apiResourceCmd)
}
