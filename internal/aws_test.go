package internal

import (
	"fmt"
	"testing"
)

func awsMocking() AWSParams {
	awsConn := NewAWS("zent-dev", "ap-northeast-2")
	awsConn.SetClusters([]string{"arn:aws:ecs:ap-northeast-2:767397666569:cluster/api-dev-Cluster-ocHE920QCXDr"})
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
	header, serviceDetails, err := conn.GetECSService()

	fmt.Println(header)
	fmt.Println(serviceDetails)
	fmt.Println(err)

}
