package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_ContentDelivery_Authentication_Parameter - This container is used for
// CDN content authentication service.
type SoftLayer_Container_Network_ContentDelivery_Authentication_Parameter struct {

	// Token - no documentation
	Token string `json:"token"`

	// CdnAccountName - no documentation
	CdnAccountName string `json:"cdnAccountName"`

	// ClientIp - no documentation
	ClientIp string `json:"clientIp"`

	// Referrer - no documentation
	Referrer string `json:"referrer"`

	// SourceUrl - no documentation
	SourceUrl string `json:"sourceUrl"`
}

func (softlayer_container_network_contentdelivery_authentication_parameter *SoftLayer_Container_Network_ContentDelivery_Authentication_Parameter) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Authentication_Parameter"
}
