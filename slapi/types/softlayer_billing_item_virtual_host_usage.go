package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Virtual_Host_Usage - The SoftLayer_Billing_Item_Virtual_Host_Usage data type
// contains general information relating to a single SoftLayer billing item for virtual machine peak
// usage
type SoftLayer_Billing_Item_Virtual_Host_Usage struct {

	// Resource - The resource for a peak virtual machine usage billing item.
	Resource *SoftLayer_Hardware `json:"resource"`

	// ResourceTableId - The resource (unique identifier) for a server billing item.
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_billing_item_virtual_host_usage *SoftLayer_Billing_Item_Virtual_Host_Usage) String() string {
	return "SoftLayer_Billing_Item_Virtual_Host_Usage"
}
