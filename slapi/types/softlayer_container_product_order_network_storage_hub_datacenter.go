package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_Storage_Hub_Datacenter - This class is used to contain a
// datacenter location and its associated active usage rate prices for object storage ordering.
type SoftLayer_Container_Product_Order_Network_Storage_Hub_Datacenter struct {

	// Location - The datacenter location where object storage is available.
	Location *SoftLayer_Location `json:"location,omitempty"`

	// UsageRatePrices - no documentation
	UsageRatePrices []*SoftLayer_Product_Item_Price `json:"usageRatePrices,omitempty"`
}

func (softlayer_container_product_order_network_storage_hub_datacenter *SoftLayer_Container_Product_Order_Network_Storage_Hub_Datacenter) String() string {
	return "SoftLayer_Container_Product_Order_Network_Storage_Hub_Datacenter"
}
