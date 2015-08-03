package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Storage_Group_Partition - A storage group partition container used
// for a hardware server order. This object describes the partitions for a single storage group that
// can be added to an order container.
type SoftLayer_Container_Product_Order_Storage_Group_Partition struct {

	// Size - no documentation
	Size float64 `json:"size"`

	// IsGrow - no documentation
	IsGrow bool `json:"isGrow"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_container_product_order_storage_group_partition *SoftLayer_Container_Product_Order_Storage_Group_Partition) String() string {
	return "SoftLayer_Container_Product_Order_Storage_Group_Partition"
}
