package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_ContentDelivery_OriginPull_Mapping - SoftLayer's CDN allows for multiple
// origin pull domains and records. This container holds the origin pull configuration details. CDN
// currently supports origin pull method for content.
type SoftLayer_Container_Network_ContentDelivery_OriginPull_Mapping struct {

	// Cname - no documentation
	Cname string `json:"cname"`

	// Id - The unique identifier of an origin pull configuration
	Id string `json:"id"`

	// IsSecureContent - This indicates if an origin pull mapping is for the secure content or not.
	IsSecureContent bool `json:"isSecureContent"`

	// MediaType - The type of a media supported by Supported media types are: and
	MediaType string `json:"mediaType"`

	// OriginUrl - The URL of a origin server. A URL can contain a directory path.
	OriginUrl string `json:"originUrl"`
}

func (softlayer_container_network_contentdelivery_originpull_mapping *SoftLayer_Container_Network_ContentDelivery_OriginPull_Mapping) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_OriginPull_Mapping"
}
