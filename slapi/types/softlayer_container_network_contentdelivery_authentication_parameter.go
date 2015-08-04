package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_ContentDelivery_Authentication_Parameter - This container is used for
// CDN content authentication service.
type SoftLayer_Container_Network_ContentDelivery_Authentication_Parameter struct {

	// CdnAccountName - no documentation
	CdnAccountName string `json:"cdnAccountName,omitempty"`

	// ClientIp - no documentation
	ClientIp string `json:"clientIp,omitempty"`

	// Referrer - no documentation
	Referrer string `json:"referrer,omitempty"`

	// SourceUrl - no documentation
	SourceUrl string `json:"sourceUrl,omitempty"`

	// Token - no documentation
	Token string `json:"token,omitempty"`
}

func (softlayer_container_network_contentdelivery_authentication_parameter *SoftLayer_Container_Network_ContentDelivery_Authentication_Parameter) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Authentication_Parameter"
}
