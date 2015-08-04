package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_ContentDelivery_Authentication_ServiceEndpoint - CDN supports the
// content authentication service. With the content authentication service, customers can control
// access to their contents. There are several scenarios where this authentication capability could be
// useful. Websites can prevent other rogue websites from linking to their videos. Content owners can
// prevent users from passing around http links, thus forcing them to login to view contents. It is
// also possible to authenticate via the client IP address. Referrer information is also checked if
// provided by a client's browser. servers will invoke a web service method to validate a content
// authentication token. CDN uses the default authentication web service provided by SoftLayer to
// validate a token. A customer can use their own implementation of the token authentication web
// service by using
// [[SoftLayer_Network_ContentDelivery_Account::setAuthenticationServiceEndpoint|setAuthenticationServiceEndpoint]]
// method. This container class holds the token validation web service endpoint information. CDN
// supports 3 different protocols: (streaming Flash), and MMS (streaming Windows Media)
type SoftLayer_Container_Network_ContentDelivery_Authentication_ServiceEndpoint struct {

	// Protocol - The protocol that the will be used for. This can be or
	Protocol string `json:"protocol,omitempty"`

	// Endpoint - The authentication web service endpoint that CDN servers will use to validate a token
	Endpoint string `json:"endpoint,omitempty"`
}

func (softlayer_container_network_contentdelivery_authentication_serviceendpoint *SoftLayer_Container_Network_ContentDelivery_Authentication_ServiceEndpoint) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Authentication_ServiceEndpoint"
}
