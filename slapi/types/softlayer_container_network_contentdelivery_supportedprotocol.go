package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_ContentDelivery_SupportedProtocol - SoftLayer's CDN content delivery
// network allows for multiple types of media hosting in addition to traditional hosting. Each of these
// media types are accessible form a different
// SoftLayer_Container_Network_ContentDelivery_SupportedProtocol holds details about CDN supported
// media types and their associated URLs. CDN media URLs follow the standard .cdn.softlayer.net Flash
// streaming, Windows Media streaming and protocols are supported: Flash streaming:
type SoftLayer_Container_Network_ContentDelivery_SupportedProtocol struct {

	// Host - The host name related to CDN supported media, and is represented in the hostname portion of a
	// CDN
	Host string `json:"host,omitempty"`

	// MediaType - no documentation
	MediaType string `json:"mediaType,omitempty"`

	// Platform - The platform name. It's a friendly name of media type.
	Platform string `json:"platform,omitempty"`

	// Protocol - The media protocol supported by This represents the media portion of a CDN Currently
	// supported protocols are: rtmp, mms and http
	Protocol string `json:"protocol,omitempty"`
}

func (softlayer_container_network_contentdelivery_supportedprotocol *SoftLayer_Container_Network_ContentDelivery_SupportedProtocol) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_SupportedProtocol"
}
