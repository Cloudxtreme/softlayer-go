package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Storage_Modification - The
// SoftLayer_Container_Product_Order_Network_Storage_Modification datatype has everything required to
// place a modification to an existing StorageLayer account with SoftLayer. Modifications, at present
// time, include upgrade and downgrades only. The ''volumeId'' property must be set to the network
// storage volume id to be upgraded. Once populated send this container to the
// [[SoftLayer_Product_Order::placeOrder]] method. The ''packageId'' property passed in for CloudLayer
// storage accounts must be set to 0 (zero) and the ''quantity'' property must be set to 1. The
// location does not have to be set. Please use the [[SoftLayer_Product_Package]] service to retrieve a
// list of CloudLayer items. When upgrading CloudLayer storage service from a metered plan (pay as you
// go) to a non-metered plan, make sure the chosen plan's storage allotment has enough space to cover
// the current usage. If the chosen plan's usage allotment is less than the CloudLayer storage's usage
// the order will be rejected.
type SoftLayer_Container_Product_Order_Network_Storage_Modification struct {

	// VolumeId - no documentation
	VolumeId int `json:"volumeId"`
}

func (softlayer_container_product_order_network_storage_modification *SoftLayer_Container_Product_Order_Network_Storage_Modification) String() string {
	return "SoftLayer_Container_Product_Order_Network_Storage_Modification"
}
