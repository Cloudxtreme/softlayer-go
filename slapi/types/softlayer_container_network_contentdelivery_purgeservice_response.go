package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_ContentDelivery_PurgeService_Response - This container holds information
// on a purge request. [[SoftLayer_Network_ContentDelivery_Account::purgeCache|Purge method]] for more
// details. Status code can be or code is returned when a URL is malformed or does not belong to
// customer. is returned in case there was an internal error.
type SoftLayer_Container_Network_ContentDelivery_PurgeService_Response struct {

	// StatusCode - A status code indicates whether your request was successful or not
	StatusCode string `json:"statusCode"`

	// Url - no documentation
	Url string `json:"url"`
}

func (softlayer_container_network_contentdelivery_purgeservice_response *SoftLayer_Container_Network_ContentDelivery_PurgeService_Response) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_PurgeService_Response"
}
