package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Storage_Group_Order - Single storage group(array) used for a hardware server
// order. If a raid configuration is required this object will describe a single array that will be
// configured on the server. If the server requires more than one array, a storage group will need to
// be created for each array.
type SoftLayer_Configuration_Storage_Group_Order struct {

	// ArrayNumber - <nil>
	ArrayNumber int `json:"arrayNumber"`

	// ArraySize - <nil>
	ArraySize float64 `json:"arraySize"`

	// ArrayType - no documentation
	ArrayType *SoftLayer_Configuration_Storage_Group_Array_Type `json:"arrayType"`

	// ArrayTypeId - <nil>
	ArrayTypeId int `json:"arrayTypeId"`

	// BillingOrderItem - no documentation
	BillingOrderItem *SoftLayer_Billing_Order_Item `json:"billingOrderItem"`

	// BillingOrderItemId - <nil>
	BillingOrderItemId int `json:"billingOrderItemId"`

	// HardDrives - <nil>
	HardDrives []int `json:"hardDrives"`

	// HotSpareDrives - <nil>
	HotSpareDrives []int `json:"hotSpareDrives"`

	// PartitionData - <nil>
	PartitionData string `json:"partitionData"`
}

func (softlayer_configuration_storage_group_order *SoftLayer_Configuration_Storage_Group_Order) String() string {
	return "SoftLayer_Configuration_Storage_Group_Order"
}
