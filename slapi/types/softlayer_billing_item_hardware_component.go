package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Hardware_Component - The SoftLayer_Billing_Item_Hardware data type contains
// general information relating to a single SoftLayer billing item for hardware components.
type SoftLayer_Billing_Item_Hardware_Component struct {

	// Resource - The hardware component that this billing item points to.
	Resource []*SoftLayer_Hardware_Component `json:"resource"`

	// ResourceCount - A count of the hardware component that this billing item points to.
	ResourceCount uint64 `json:"resourceCount"`

	// ResourceTableId - The resource (unique identifier) for a server billing item.
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_billing_item_hardware_component *SoftLayer_Billing_Item_Hardware_Component) String() string {
	return "SoftLayer_Billing_Item_Hardware_Component"
}
