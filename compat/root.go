package compat

import (
	"github.com/go-co-op/gocron"
	"github.com/oskardotglobal/uptimekuma-cli/util"
)

type Node interface {
	GetName() string
	ShouldReportStatus() bool
}

func ReportNodes(scheduler *gocron.Scheduler) {
	ReportDocker(scheduler)
}

func ReportStatusForNode(node Node) {
	if node.ShouldReportStatus() {
		util.ReportStatus("nodes." + node.GetName())
	}
}
