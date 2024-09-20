package internal

import (
	"testing"
)

func awsMocking() AWSParams {
	awsConn := NewAWS("zent-dev", "ap-northeast-2")
	awsConn.SetClusters([]string{"arn:aws:ecs:ap-northeast-2:767397666569:cluster/devops-common-Cluster-2Zb36ZXZRqbI"})
	return awsConn
}

func Test_GetECSClusterDetails(t *testing.T) {

	conn := awsMocking()
	_, clusterDetails, err := conn.GetECSClusterDetails()
	if err != nil {
		t.Error(err)
	}

	if len(clusterDetails) == 0 {
		t.Error("No cluster details")
	}
}

func Test_GetECSService(t *testing.T) {
	conn := awsMocking()
	_, serviceDetails, err := conn.GetECSService()

	if err != nil {
		t.Error(err)
	}

	if len(serviceDetails) == 0 {
		t.Error("No cluster details")
	}

}

func Test_GetECSContainer(t *testing.T) {
	conn := awsMocking()
	_, containerDetails, err := conn.GetECSContainers()

	if err != nil {
		t.Error(err)
	}

	if len(containerDetails) == 0 {
		t.Error("No cluster details")
	}
}
