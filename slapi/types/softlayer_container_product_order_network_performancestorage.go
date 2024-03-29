package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_PerformanceStorage - This is the base data type for
// Performance storage order containers. If you wish to place an order you must not use this class and
// instead use the appropriate child container for the type of storage you would like to order:
// [[SoftLayer_Container_Product_Order_Network_PerformanceStorage_Nfs]] for File and
// [[SoftLayer_Container_Product_Order_Network_PerformanceStorage_Iscsi]] for Block storage.
type SoftLayer_Container_Product_Order_Network_PerformanceStorage struct {
}

func (softlayer_container_product_order_network_performancestorage *SoftLayer_Container_Product_Order_Network_PerformanceStorage) String() string {
	return "SoftLayer_Container_Product_Order_Network_PerformanceStorage"
}
