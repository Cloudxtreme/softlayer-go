package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Storage_Iscsi_SnapshotSpace - This is the datatype that
// needs to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything
// required to place an Snapshot Space order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_Storage_Iscsi_SnapshotSpace struct {

	// VolumeId - the [[SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3]] Id.
	VolumeId int `json:"volumeId"`
}

func (softlayer_container_product_order_network_storage_iscsi_snapshotspace *SoftLayer_Container_Product_Order_Network_Storage_Iscsi_SnapshotSpace) String() string {
	return "SoftLayer_Container_Product_Order_Network_Storage_Iscsi_SnapshotSpace"
}
