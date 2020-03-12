package azuredisk

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"github.com/giantswarm/azure-disk-mitigator/service/controller/key"
	"github.com/giantswarm/microerror"
	"regexp"
)

var (
	detachErrorRegex = regexp.MustCompile("AttachVolume\\.Attach failed for volume \"(.*?)\" : disk\\(\\/subscriptions\\/(.*?)\\/resourceGroups\\/(.*?)\\/providers\\/Microsoft\\.Compute\\/disks\\/(.*?)\\) already attached to node\\(\\/subscriptions\\/(.*?)\\/resourceGroups\\/(.*?)\\/providers\\/Microsoft\\.Compute\\/virtualMachineScaleSets\\/(.*?)\\/virtualMachines\\/(.*?)\\), could not be attached to node\\((.*?)\\)")
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	event, err := key.ToEvent(obj)
	if err != nil {
		return microerror.Mask(err)
	}

	if event.Type != corev1.EventTypeWarning || event.Reason != "FailedAttachVolume" {
		return nil
	}

	r.logger.LogCtx(ctx, "message", event.Message)
	match := detachErrorRegex.FindStringSubmatch(event.Message)
	pvcName := match[1]
	r.logger.LogCtx(ctx, "pvc", pvcName)
	// subscriptionId := match[2]
	// resourceGroup := match[3]
	// diskName := match[4]
	// subscriptionId := match[5]
	// resourceGroup := match[6]
	// vmssName := match[7]
	// vmssInstanceName := match[8]
	// nodeName := match[9]

	return nil
}
