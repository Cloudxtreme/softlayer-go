package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Container_Product_Order_Storage_Group - A single storage group container used for a
// hardware server order. This object describes a single storage group that can be added to an order
// container.
type SoftLayer_Container_Product_Order_Storage_Group struct {

	// HotSpareDrives - If an array should be protected by an hotspare, the drive index of the hotspares
	// should be here. If a drive is a hotspare for all arrays then a separate storage group with array
	// type should be used
	HotSpareDrives []int `json:"hotSpareDrives,omitempty"`

	// PartitionTemplateId - The id for a [[SoftLayer_Hardware_Component_Partition_Template]] object, which
	// will determine the partitions to add to the storage group. If this storage group is not a primary
	// storage group, then this will not be used.
	PartitionTemplateId int `json:"partitionTemplateId,omitempty"`

	// Partitions - Defines the partitions for the storage group. If this storage group is not a secondary
	// storage group, then this will not be used.
	Partitions []*SoftLayer_Container_Product_Order_Storage_Group_Partition `json:"partitions,omitempty"`

	// ArraySize - Size of the array in gigabytes. Must be within limitations of the smallest drive
	// assigned to the storage group and the storage group type.
	ArraySize slapi.Float64 `json:"arraySize,omitempty"`

	// ArrayTypeId - The array type id from a [[SoftLayer_Configuration_Storage_Group_Array_Type]] object.
	ArrayTypeId int `json:"arrayTypeId,omitempty"`

	// HardDrives - Integer array of drive indexes to use in the storage group.
	HardDrives []int `json:"hardDrives,omitempty"`
}

func (softlayer_container_product_order_storage_group *SoftLayer_Container_Product_Order_Storage_Group) String() string {
	return "SoftLayer_Container_Product_Order_Storage_Group"
}
