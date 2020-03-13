[![CircleCI](https://circleci.com/gh/giantswarm/azure-disk-mitigator-app.svg?&style=shield)](https://circleci.com/gh/giantswarm/azure-disk-mitigator-app) [![Docker Repository on Quay](https://quay.io/repository/giantswarm/azure-disk-mitigator-app/status "Docker Repository on Quay")](https://quay.io/repository/giantswarm/azure-disk-mitigator-app)

# azure-disk-mitigator-app

Azure Disk mitigator is an operator for mitigating Azure Disk attach/detach issues which occur in
Kubernetes clusters.

When a PVC tries to move from one node to another, detachment of Azure disk from VMSS instance
can fail and a PVC will get stuck. This operator tries to detach the disk when such event occurs.