<!--

    TODO:

    - Add the project to the CircleCI:
      https://circleci.com/setup-project/gh/giantswarm/REPOSITORY_NAME

    - Import RELEASE_TOKEN variable from template repository for the builds:
      https://circleci.com/gh/giantswarm/REPOSITORY_NAME/edit#env-vars

    - Change the badge (with style=shield):
      https://circleci.com/gh/giantswarm/REPOSITORY_NAME/edit#badges
      If this is a private repository token with scope `status` will be needed.

    - Run `devctl replace -i "REPOSITORY_NAME" "$(basename $(git rev-parse --show-toplevel))" *.md`
      and commit your changes.

-->
[![CircleCI](https://circleci.com/gh/giantswarm/azure-disk-mitigator-app.svg?&style=shield)](https://circleci.com/gh/giantswarm/azure-disk-mitigator-app) [![Docker Repository on Quay](https://quay.io/repository/giantswarm/azure-disk-mitigator-app/status "Docker Repository on Quay")](https://quay.io/repository/giantswarm/azure-disk-mitigator-app)

# REPOSITORY_NAME

This is a template repository containing files for a giantswarm
operator repository.

To use it just hit `Use this template` button or [this
link][generate].

After creating your repository replace all instances of
`azure-disk-mitigator-app` in this code base with your new repository name.
Also rename `helm/azure-disk-mitigator-app` directory.

[generate]: https://github.com/giantswarm/azure-disk-mitigator-app/generate
