package internal

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func (ap AWSParams) GetECSCluster() ([]string, error) {

	res, err := ap.ecsClient.ListClusters(context.TODO(), &ecs.ListClustersInput{})
	if err != nil {
		return nil, err
	}

	clusterNames := make([]string, len(res.ClusterArns))
	copy(clusterNames, res.ClusterArns)

	return clusterNames, nil
}
