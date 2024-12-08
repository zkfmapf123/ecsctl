package internal

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/zkfmapf123/dobbyssm/utils"
)

type AWSParams struct {
	cluster   []string
	ecsClient *ecs.Client
}

func NewAWS(profile string, region string) AWSParams {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)
	if err != nil {
		utils.PanicRed(err)
	}

	return AWSParams{
		ecsClient: ecs.NewFromConfig(cfg),
	}
}

func (ap *AWSParams) SetClusters(clusters []string) {

	if len(clusters) == 0 {
		cs, err := ap.GetECSCluster()
		if err != nil {
			utils.PanicRed(err)
		}

		ap.cluster = cs
	} else {
		ap.cluster = clusters
	}
}
