package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_PerformanceStorage_Iscsi - This is the datatype that needs
// to be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything
// required to place an order for iSCSI (Block) Performance Storage
type SoftLayer_Container_Product_Order_Network_PerformanceStorage_Iscsi struct {

	// OsFormatType - OS Type to be used when formatting the storage space, this should match the OS type
	// that will be connecting to the The only required property its the keyName of the OS type.
	OsFormatType *SoftLayer_Network_Storage_Iscsi_OS_Type `json:"osFormatType"`
}

func (softlayer_container_product_order_network_performancestorage_iscsi *SoftLayer_Container_Product_Order_Network_PerformanceStorage_Iscsi) String() string {
	return "SoftLayer_Container_Product_Order_Network_PerformanceStorage_Iscsi"
}
