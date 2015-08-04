package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Configuration_Storage_Group_Order - Single storage group(array) used for a hardware server
// order. If a raid configuration is required this object will describe a single array that will be
// configured on the server. If the server requires more than one array, a storage group will need to
// be created for each array.
type SoftLayer_Configuration_Storage_Group_Order struct {

	// ArraySize - <nil>
	ArraySize slapi.Float64 `json:"arraySize,omitempty"`

	// BillingOrderItemId - <nil>
	BillingOrderItemId int `json:"billingOrderItemId,omitempty"`

	// ArrayTypeId - <nil>
	ArrayTypeId int `json:"arrayTypeId,omitempty"`

	// HardDrives - <nil>
	HardDrives []int `json:"hardDrives,omitempty"`

	// HotSpareDrives - <nil>
	HotSpareDrives []int `json:"hotSpareDrives,omitempty"`

	// PartitionData - <nil>
	PartitionData string `json:"partitionData,omitempty"`

	// ArrayNumber - <nil>
	ArrayNumber int `json:"arrayNumber,omitempty"`
}

func (softlayer_configuration_storage_group_order *SoftLayer_Configuration_Storage_Group_Order) String() string {
	return "SoftLayer_Configuration_Storage_Group_Order"
}

// SoftLayer_Configuration_Storage_Group_Order_Extended is SoftLayer_Configuration_Storage_Group_Order with all maskable types.
type SoftLayer_Configuration_Storage_Group_Order_Extended struct {
	SoftLayer_Configuration_Storage_Group_Order

	// ArrayType - no documentation
	ArrayType *SoftLayer_Configuration_Storage_Group_Array_Type `json:"arrayType,omitempty"`

	// BillingOrderItem - no documentation
	BillingOrderItem *SoftLayer_Billing_Order_Item `json:"billingOrderItem,omitempty"`
}

func (softlayer_configuration_storage_group_order *SoftLayer_Configuration_Storage_Group_Order_Extended) String() string {
	return "SoftLayer_Configuration_Storage_Group_Order"
}
