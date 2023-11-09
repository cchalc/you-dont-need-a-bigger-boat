package main

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

func ListAllClusters() {
	w := databricks.Must(databricks.NewWorkspaceClient())
	all, err := w.Clusters.ListAll(context.Background(), compute.ListClustersRequest{})
	if err != nil {
		panic(err)
	}
	for _, c := range all {
		println(c.ClusterName)
	}
}

func CreateLTSCluster() {
	const clusterName = "cjc-gosdk"
	const autoTerminationMinutes = 40
	const numWorkers = 2

	w := databricks.Must(databricks.NewWorkspaceClient())
	ctx := context.Background()

	// get the full list of available spark versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)

	if err != nil {
		panic(err)
	}

	// choose the latest long term support
	latestLTS, err := sparkVersions.Select(compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})

	if err != nil {
		panic(err)
	}

	nodeTypes, err := w.Clusters.ListNodeTypes(ctx)

	if err != nil {
		panic(err)
	}

	smallestWithLocalDisk, err := nodeTypes.Smallest(compute.NodeTypeRequest{
		LocalDisk: true,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("Now attempting to create the cluster, please wait...")

	runningCluster, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latestLTS,
		NodeTypeId:             smallestWithLocalDisk,
		AutoterminationMinutes: autoTerminationMinutes,
		NumWorkers:             numWorkers,
	})
	if err != nil {
		panic(err)
	}
	switch runningCluster.State {
	case compute.StateRunning:
		fmt.Printf("The cluster is now ready at %s#setting/clusters/%s/configuration\n",
			w.Config.Host,
			runningCluster.ClusterId,
		)
	default:
		fmt.Printf("Cluster is not running or failed to create. %s", runningCluster.StateMessage)
	}
}

func main() {
	ListAllClusters()
	CreateLTSCluster()
}
