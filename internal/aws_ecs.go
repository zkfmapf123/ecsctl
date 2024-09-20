package internal

import (
	"context"
	"errors"
	"strconv"
	"strings"

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

func (ap AWSParams) getECSServiceDetails(cluster string) (*ecs.DescribeServicesOutput, error) {

	res, err := ap.ecsClient.ListServices(context.TODO(), &ecs.ListServicesInput{
		Cluster: &cluster,
	})

	if err != nil {
		return nil, err
	}

	resService, err := ap.ecsClient.DescribeServices(context.TODO(), &ecs.DescribeServicesInput{
		Cluster:  &cluster,
		Services: res.ServiceArns,
	})

	if err != nil {
		return nil, err
	}

	return resService, nil
}

func (ap AWSParams) getECSTasks(cluster, serviceName string) ([]string, error) {

	res, err := ap.ecsClient.ListTasks(context.TODO(), &ecs.ListTasksInput{
		Cluster:     &cluster,
		ServiceName: &serviceName,
	})

	if err != nil {
		return nil, err
	}

	return res.TaskArns, nil
}

func (ap AWSParams) GetECSService() ([]string, [][]string, error) {

	if len(ap.cluster) == 0 {
		return nil, nil, errors.New("No cluster found Use ecsns")
	}

	values := [][]string{}
	for _, cluster := range ap.cluster {

		resService, err := ap.getECSServiceDetails(cluster)
		if err != nil {
			return nil, nil, err
		}

		for _, v := range resService.Services {

			revision := *v.TaskDefinition
			revisionArr := strings.Split(revision, "/")

			values = append(
				values,
				[]string{
					*v.ServiceName,
					*v.Status,
					revisionArr[len(revisionArr)-1],
					strconv.Itoa(int(v.RunningCount)),
					strconv.Itoa(int(v.PendingCount)),
					strconv.Itoa(int(v.DesiredCount)),
					v.CreatedAt.String(),
					strings.Join(v.NetworkConfiguration.AwsvpcConfiguration.Subnets, " | "),
					string(v.NetworkConfiguration.AwsvpcConfiguration.AssignPublicIp),
					strconv.FormatBool(v.EnableExecuteCommand),
				},
			)
		}
	}

	return []string{
		"Service Name",
		"Status",
		"Revision",
		"Running Count",
		"Pending Count",
		"Desired Count",
		"Created At",
		"Network : Subnest",
		"Network : Assign Public Ip",
		"exec",
	}, values, nil
}

func (ap AWSParams) GetECSContainers() ([]string, [][]string, error) {

	if len(ap.cluster) == 0 {
		return nil, nil, errors.New("No cluster found Use ecsns")
	}

	values := [][]string{}
	for _, cluster := range ap.cluster {

		// serivce
		serviceRes, err := ap.getECSServiceDetails(cluster)
		if err != nil {
			return nil, nil, err
		}

		// Task
		for _, svc := range serviceRes.Services {

			taskRes, err := ap.getECSTasks(cluster, *svc.ServiceName)
			if err != nil {
				return nil, nil, err
			}

			// containers
			containerRes, err := ap.ecsClient.DescribeTasks(context.TODO(), &ecs.DescribeTasksInput{
				Cluster: &cluster,
				Tasks:   taskRes,
			})
			if err != nil {
				return nil, nil, err
			}

			for _, t := range containerRes.Tasks {

				// clusterName
				_cluster := strings.Split(cluster, "/")
				clusterName := _cluster[len(_cluster)-1]

				// containerArn
				cArn := *t.TaskArn
				containerArn := strings.Split(cArn, "/")

				for _, container := range t.Containers {

					image := ""
					if container.Image != nil {
						image = *container.Image
					}

					values = append(
						values,
						[]string{
							clusterName,
							containerArn[len(containerArn)-1],
							*container.Name,
							image,
							*container.LastStatus,
							t.StartedAt.String(),
							*t.Cpu,
							*t.Memory,
							string(t.LaunchType),
						},
					)

				}
			}
		}
	}

	return []string{
		"Cluster Name",
		"Task Id",
		"Container Name",
		"Image",
		"Status",
		"Task Started At",
		"CPU",
		"Memory",
		"Launch Type",
	}, values, nil
}
