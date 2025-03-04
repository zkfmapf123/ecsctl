package internal

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/zkfmapf123/dobbyssm/utils"
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

func (ap AWSParams) getECSServiceDetails(cluster string) ([]types.Service, error) {

	var serviceArns []string
	var nextToken *string

	for {
		res, err := ap.ecsClient.ListServices(context.TODO(), &ecs.ListServicesInput{
			Cluster:   &cluster,
			NextToken: nextToken,
		})

		if err != nil {
			return nil, err
		}

		serviceArns = append(serviceArns, res.ServiceArns...)
		if res.NextToken == nil {
			break
		}

		if res.NextToken == nil {
			break
		}

		nextToken = res.NextToken
	}

	resServiceOutput := []types.Service{}

	for i := 0; i < len(serviceArns); i += 10 {
		end := i + 10
		if end > len(serviceArns) {
			end = len(serviceArns)
		}

		resService, err := ap.ecsClient.DescribeServices(context.TODO(), &ecs.DescribeServicesInput{
			Cluster:  &cluster,
			Services: serviceArns[i:end],
		})

		if err != nil {
			if strings.Contains(err.Error(), utils.EXCEPTION_EMPTY_SERVICE) {
				return nil, errors.New(utils.EXCEPTION_EMPTY_SERVICE)
			}

			return nil, err

		}

		resServiceOutput = append(resServiceOutput, resService.Services...)
	}

	return resServiceOutput, nil
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

			if strings.Contains(err.Error(), utils.EXCEPTION_EMPTY_SERVICE) {
				continue
			}

			return nil, nil, err
		}

		for _, v := range resService {

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

			if strings.Contains(err.Error(), utils.EXCEPTION_EMPTY_SERVICE) {
				continue
			}

			return nil, nil, err
		}

		// Task
		for _, svc := range serviceRes {

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

					values = append(
						values,
						[]string{
							clusterName,
							containerArn[len(containerArn)-1],
							*container.Name,
							// image,
							*container.LastStatus,
							t.StartedAt.String(),
							*t.Cpu,
							*t.Memory,
							string(t.LaunchType),
							strconv.FormatBool(t.EnableExecuteCommand),
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
		// "Image",
		"Status",
		"Task Started At",
		"CPU",
		"Memory",
		"Launch Type",
		"Exec",
	}, values, nil
}
