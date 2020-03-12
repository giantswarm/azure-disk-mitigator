package controller

import (
	// If your operator watches a CRD import it here.
	// "github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/giantswarm/azure-disk-mitigator/pkg/project"
)

type EventConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

type Event struct {
	*controller.Controller
}

func NewEvent(config EventConfig) (*Event, error) {
	var err error

	resourceSets, err := newEventResourceSets(config)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var operatorkitController *controller.Controller
	{
		c := controller.Config{
			// If your operator watches a CRD add it here.
			// CRD:       v1alpha1.NewAppCRD(),
			K8sClient:    config.K8sClient,
			Logger:       config.Logger,
			ResourceSets: resourceSets,
			NewRuntimeObjectFunc: func() runtime.Object {
				return new(corev1.Event)
			},

			// Name is used to compute finalizer names. This here results in something
			// like operatorkit.giantswarm.io/azure-disk-mitigator-event-controller.
			Name: project.Name() + "-event-controller",
		}

		operatorkitController, err = controller.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	c := &Event{
		Controller: operatorkitController,
	}

	return c, nil
}

func newEventResourceSets(config EventConfig) ([]*controller.ResourceSet, error) {
	var err error

	var resourceSet *controller.ResourceSet
	{
		c := eventResourceSetConfig{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		resourceSet, err = newEventResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resourceSets := []*controller.ResourceSet{
		resourceSet,
	}

	return resourceSets, nil
}
