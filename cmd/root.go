package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/dobbyssm/internal"
	"github.com/zkfmapf123/dobbyssm/utils"
)

const (
	_defaultProfile = "default"
	_defaultRegion  = "ap-northeast-2"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ecsctl",
		Short: "ecsctl is interactive CLI for AWS ECS",
		Long:  "ecsctl is interactive CLI for AWS ECS",
		Run: func(cmd *cobra.Command, args []string) {
			creds, err := internal.GetCredentialFile()
			t := utils.NewTerminal("")
			t.Clear()

			if err != nil {
				utils.PanicRed(err)
			}

			fmt.Printf("Profile : %s\nRegion : %s\n", creds.Profile, creds.Region)

			credClusters := make([][]string, len(creds.Clusters))
			for i, cluster := range creds.Clusters {
				credClusters[i] = []string{cluster}
			}

			t.TableWriter([]string{"Cluster"}, credClusters)

		},
	}
)

func MustExecute() {
	if err := rootCmd.Execute(); err != nil {
		utils.PanicRed(err)
	}
}

type AWSCredentials struct {
	Profile  string   `json:"profile"`
	Region   string   `json:"region"`
	Clusters []string `json:"clusters"`
}

func initConfig() {

	creds, err := internal.GetCredentialFile()
	if err != nil {
		utils.NoticeGreen(err)

		err := internal.SetCredentialFile(
			_defaultProfile, _defaultRegion, []string{},
		)

		if err != nil {
			utils.PanicRed(err)
		}
	}

	if creds.Profile == "" || creds.Region == "" {
		utils.NoticeGreen(errors.New("Please set your aws profile and region"))
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
