package compat

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-co-op/gocron"
	"github.com/oskardotglobal/uptimekuma-cli/util"
	"github.com/spf13/cobra"
)

type DockerNode struct {
	Name    string
	ID      string
	Client  client.APIClient
	Context context.Context
}

// GetName this implicitly implements the Node inteface for DockerNode
// no need for an "implements" keyword
func (node DockerNode) GetName() string {
	if node.Name != "" {
		return node.Name
	}

	return node.ID
}

func (node DockerNode) ShouldReportStatus() bool {
	status, err := node.Client.ContainerInspect(node.Context, node.ID)
	cobra.CheckErr(err)

	return status.State.Running
}

func GetDockerContainers(cli client.APIClient) []DockerNode {
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	cobra.CheckErr(err)

	var ret []DockerNode

	for _, container := range containers {
		id := container.ID[:10]
		inspect, err := cli.ContainerInspect(ctx, id)

		cobra.CheckErr(err)

		ret = append(ret, DockerNode{
			Name:    inspect.Name,
			ID:      id,
			Client:  cli,
			Context: ctx,
		})
	}

	return ret
}

func ReportDocker(scheduler *gocron.Scheduler) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		util.Warn(err, "Couldn't connect to docker, skipping... - "+err.Error())
		return
	}

	for _, container := range GetDockerContainers(cli) {
		_, err := scheduler.Every(1).Minute().Do(ReportStatusForNode, container)
		util.CheckErrorWithMsg(err, "Couldn't schedule task for container "+container.GetName())
	}

	defer cli.Close()
}
