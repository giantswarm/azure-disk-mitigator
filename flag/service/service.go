package service

import (
	"github.com/giantswarm/operatorkit/flag/service/kubernetes"

	"github.com/giantswarm/azure-disk-mitigator-app/flag/service/azure"
)

// Service is an intermediate data structure for command line configuration flags.
type Service struct {
	Azure      azure.Azure
	Kubernetes kubernetes.Kubernetes
}
