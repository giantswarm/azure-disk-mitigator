package project

var (
	bundleVersion = "0.0.1-dev"
	description   = "The azure-disk-mitigator-app does something."
	gitSHA        = "n/a"
	name          = "azure-disk-mitigator-app"
	source        = "https://github.com/giantswarm/azure-disk-mitigator-app"
	version       = "n/a"
)

func BundleVersion() string {
	return bundleVersion
}

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
