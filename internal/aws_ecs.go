package internal

import (
	"context"
	"errors"
	"strconv"

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

type ClusterDetails struct {
	ClusterName                       string
	Status                            string
	RunningTasksCount                 int32
	PendingTasksCount                 int32
	ActiveServicesCount               int32
	RegisteredContainerInstancesCount int32
}

func (ap AWSParams) GetECSClusterDetails() ([]string, [][]string, error) {

	clusterArns, err := ap.GetECSCluster()
	if err != nil {
		return nil, nil, err
	}

	res, err := ap.ecsClient.DescribeClusters(context.TODO(), &ecs.DescribeClustersInput{
		Clusters: clusterArns,
	})

	if err != nil {
		return nil, nil, err
	}

	values := make([][]string, len(res.Clusters))

	for _, v := range res.Clusters {
		v := []string{
			*v.ClusterName,
			*v.Status,
			strconv.Itoa(int(v.RunningTasksCount)),
			strconv.Itoa(int(v.PendingTasksCount)),
			strconv.Itoa(int(v.ActiveServicesCount)),
			strconv.Itoa(int(v.RegisteredContainerInstancesCount)),
		}
		values = append(values, v)
	}

	// headers
	headers := []string{
		"ClusterName",
		"Status",
		"RunningTasksCount",
		"PendingTasksCount",
		"ActiveServicesCount",
		"RegisteredContainerInstancesCount",
	}

	return headers, values, err
}

func (ap AWSParams) GetECSService() ([]string, [][]string, error) {

	if len(ap.cluster) == 0 {
		return nil, nil, errors.New("No cluster found Use ecsns")
	}

	values := [][]string{}
	for _, cluster := range ap.cluster {

		res, err := ap.ecsClient.ListServices(context.TODO(), &ecs.ListServicesInput{
			Cluster: &cluster,
		})

		if err != nil {
			return nil, nil, err
		}

		resService, err := ap.ecsClient.DescribeServices(context.TODO(), &ecs.DescribeServicesInput{
			Cluster:  &cluster,
			Services: res.ServiceArns,
		})

		if err != nil {
			return nil, nil, err
		}

		for _, v := range resService.Services {
			values = append(
				values,
				[]string{
					*v.ServiceName,
					*v.Status,
					strconv.Itoa(int(v.RunningCount)),
					strconv.Itoa(int(v.PendingCount)),
					strconv.Itoa(int(v.DesiredCount)),
					v.CreatedAt.String(),
					strconv.FormatBool(v.EnableExecuteCommand),
				},
			)

			// fmt.Println("TaskDefinition : ", v.TaskDefinition)
			// fmt.Println("Events : ", v.Events)
			// fmt.Println("NetworkConfiguration : ", v.NetworkConfiguration)
			// fmt.Println("LoadBalancers : ", v.LoadBalancers)
		}
	}

	return []string{
		"Service Name",
		"Status",
		"Running Count",
		"Pending Count",
		"Desired Count",
		"Created At",
		"exec",
	}, values, nil
}
