package client

type AzureClientSetConfig struct {
	// ClientID is the ID of the Active Directory Service Principal.
	ClientID string
	// ClientSecret is the secret of the Active Directory Service Principal.
	ClientSecret string
	// EnvironmentName is the cloud environment identifier on Azure. Values can be
	// used as listed in the link below.
	//
	//     https://github.com/Azure/go-autorest/blob/ec5f4903f77ed9927ac95b19ab8e44ada64c1356/autorest/azure/environments.go#L13
	//
	EnvironmentName string
	// SubscriptionID is the ID of the Azure subscription.
	SubscriptionID string
	// TenantID is the ID of the Active Directory tenant.
	TenantID string
	// PartnerID is the ID used for the Azure Partner Program.
	PartnerID string
}
