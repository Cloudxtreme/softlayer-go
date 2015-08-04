package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl -
// SoftLayer_Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl provides specific details
// is a container which contains the cdn urls associated with an object storage account
type SoftLayer_Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl struct {

	// Datacenter - <nil>
	Datacenter string `json:"datacenter,omitempty"`

	// FlashUrl - <nil>
	FlashUrl string `json:"flashUrl,omitempty"`

	// HttpUrl - <nil>
	HttpUrl string `json:"httpUrl,omitempty"`
}

func (softlayer_container_network_storage_hub_objectstorage_contentdeliveryurl *SoftLayer_Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl) String() string {
	return "SoftLayer_Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl"
}
