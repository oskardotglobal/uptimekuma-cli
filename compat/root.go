package compat

import (
	"github.com/go-co-op/gocron"
	"github.com/oskardotglobal/uptimekuma-cli/util"
	"github.com/spf13/viper"
)

var (
	nodes []Node
)

type Node interface {
	GetName() string
	ShouldReportStatus() bool
}

func init() {
	dockerContainers := util.ArrayMap(GetDockerContainers(), func(x DockerNode) Node {
		return Node(x)
	})

	nodes = append(nodes, dockerContainers...)

	util.ArrayMap(nodes, func(x Node) string {
		util.SetNodeUrlIfEmpty(x.GetName())
		return ""
	})
}

func ReportNodes(scheduler *gocron.Scheduler, viperInstance *viper.Viper) {
	for _, node := range nodes {
		_, err := scheduler.Every(1).Minute().Do(ReportStatusForNode, node, viperInstance)
		util.CheckErrorWithMsg(err, "Couldn't schedule task for container "+node.GetName())
	}
}

func ReportStatusForNode(node Node, viperInstance *viper.Viper) {
	if node.ShouldReportStatus() {
		util.ReportStatus(viperInstance, "nodes."+node.GetName())
	}
}
