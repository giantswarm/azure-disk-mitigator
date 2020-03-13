package controller

import (
	"fmt"

	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"github.com/giantswarm/operatorkit/resource"
	"github.com/giantswarm/operatorkit/resource/wrapper/metricsresource"
	"github.com/giantswarm/operatorkit/resource/wrapper/retryresource"

	"github.com/giantswarm/azure-disk-mitigator-app/service/controller/key"
	"github.com/giantswarm/azure-disk-mitigator-app/service/controller/resource/azuredisk"
)

type eventResourceSetConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

func newEventResourceSet(config eventResourceSetConfig) (*controller.ResourceSet, error) {
	var err error

	var azureDiskResource resource.Interface
	{
		c := azuredisk.Config{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		azureDiskResource, err = azuredisk.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resources := []resource.Interface{
		azureDiskResource,
	}

	{
		c := retryresource.WrapConfig{
			Logger: config.Logger,
		}

		resources, err = retryresource.Wrap(resources, c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	{
		c := metricsresource.WrapConfig{}

		resources, err = metricsresource.Wrap(resources, c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	// handlesFunc defines which objects you want to get into your controller, e.g. which objects you want to watch.
	handlesFunc := func(obj interface{}) bool {
		cr, err := key.ToEvent(obj)
		if err != nil {
			config.Logger.Log("level", "warning", "message", fmt.Sprintf("invalid object: %s", err), "stack", fmt.Sprintf("%v", err)) // nolint: errcheck
			return false
		}

		if key.EventIsWarning(cr) && key.EventReason(cr) == "FailedAttachVolume" {
			return true
		}

		return false
	}

	var resourceSet *controller.ResourceSet
	{
		c := controller.ResourceSetConfig{
			Handles:   handlesFunc,
			Logger:    config.Logger,
			Resources: resources,
		}

		resourceSet, err = controller.NewResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return resourceSet, nil
}
