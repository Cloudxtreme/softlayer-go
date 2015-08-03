package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Storage_Enterprise - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to
// place an order for Enterprise Storage
type SoftLayer_Container_Product_Order_Network_Storage_Enterprise struct {

	// OsFormatType - This must be populated for block storage orders. This should match the OS type of the
	// host(s) that will connect to the volume. The only required property is the keyName of the OS type.
	// This property is ignored for file storage orders.
	OsFormatType *SoftLayer_Network_Storage_Iscsi_OS_Type `json:"osFormatType"`

	// OriginVolumeId - This must be populated only for replicant volume ordering. It represents the
	// identifier of the origin [[SoftLayer_Network_Storage]].
	OriginVolumeId int `json:"originVolumeId"`

	// OriginVolumeScheduleId - This must be populated only for replicant volume ordering. It represents
	// the [[SoftLayer_Network_Storage_Schedule]] that will be be used to replicate the origin
	// [[SoftLayer_Network_Storage]] volume.
	OriginVolumeScheduleId int `json:"originVolumeScheduleId"`
}

func (softlayer_container_product_order_network_storage_enterprise *SoftLayer_Container_Product_Order_Network_Storage_Enterprise) String() string {
	return "SoftLayer_Container_Product_Order_Network_Storage_Enterprise"
}
