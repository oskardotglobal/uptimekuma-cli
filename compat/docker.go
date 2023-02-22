package compat

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/oskardotglobal/uptimekuma-cli/util"
	"github.com/spf13/cobra"
)

type DockerNode struct {
	Node
	Name    string
	ID      string
	Client  client.APIClient
	Context context.Context
}

var (
	cli, cliErr = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
)

func init() {
	if cliErr != nil {
		util.Warn(cliErr, "Couldn't connect to docker, skipping... - "+cliErr.Error())
		return
	}
}

// GetName this implicitly implements the Node interface for DockerNode
// no need for an "implements" keyword
func (node DockerNode) GetName() string {
	if node.Name != "" {
		return node.Name
	}

	return node.ID
}

// ShouldReportStatus this implicitly implements the Node interface for DockerNode
func (node DockerNode) ShouldReportStatus() bool {
	status, err := node.Client.ContainerInspect(node.Context, node.ID)
	cobra.CheckErr(err)

	return status.State.Running
}

func GetDockerContainers() []DockerNode {
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
