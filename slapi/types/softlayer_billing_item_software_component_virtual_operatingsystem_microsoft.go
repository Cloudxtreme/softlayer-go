package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft - The
// SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft data type contains
// general information relating to a single SoftLayer billing item for a Microsoft operating system
// software components on virtual machines.
type SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft struct {

	// Resource - The software virtual license to which this billing item points.
	Resource *SoftLayer_Software_VirtualLicense `json:"resource"`

	// ResourceTableId - The resource (unique identifier) for a software virtual license billing item.
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_billing_item_software_component_virtual_operatingsystem_microsoft *SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft) String() string {
	return "SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft"
}
