package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace - This is the datatype
// that needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has
// everything required to place an order for Enterprise Storage Snapshot Space.
type SoftLayer_Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace struct {

	// VolumeId - The [[SoftLayer_Network_Storage]] id for which snapshot space is being ordered for.
	VolumeId int `json:"volumeId"`
}

func (softlayer_container_product_order_network_storage_enterprise_snapshotspace *SoftLayer_Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace) String() string {
	return "SoftLayer_Container_Product_Order_Network_Storage_Enterprise_SnapshotSpace"
}
