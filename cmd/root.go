package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

			fmt.Println("hello world")
		},
	}
)

func MustExecute() {
	if err := rootCmd.Execute(); err != nil {
		utils.PanicRed(err)
	}
}

type AWSCredentials struct {
	Profile string `json:"profile"`
	Region  string `json:"region"`
	Cluster string `json:"cluster"`
}

func initConfig() {

	cred, err := internal.GetCredentialFile()
	if err != nil {
		utils.NoticeGreen(err)

		// internal.SetCredentialFile()
	}

	fmt.Println(cred)

}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("profile", "p", "", `[optional] your aws profile, default is "default"`)
	rootCmd.PersistentFlags().StringP("region", "r", "", `[optional] your aws region, default is "region"`)
	rootCmd.InitDefaultVersionFlag()

	utils.MustCheckError(viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile")))
	utils.MustCheckError(viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region")))
}
