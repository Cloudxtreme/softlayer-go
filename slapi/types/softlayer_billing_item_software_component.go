package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Software_Component - The SoftLayer_Billing_Item_Hardware data type contains
// general information relating to a single SoftLayer billing item for hardware components.
type SoftLayer_Billing_Item_Software_Component struct {

	// Resource - The software component that this billing item points to.
	Resource *SoftLayer_Software_Component `json:"resource"`

	// ResourceTableId - The resource (unique identifier) for a software component billing item.
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_billing_item_software_component *SoftLayer_Billing_Item_Software_Component) String() string {
	return "SoftLayer_Billing_Item_Software_Component"
}
