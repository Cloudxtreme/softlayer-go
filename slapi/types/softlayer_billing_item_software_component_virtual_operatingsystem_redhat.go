package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat - The
// SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft data type contains
// general information relating to a single SoftLayer billing item for a Microsoft operating system
// software components on virtual machines.
type SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat struct {

	// Resource - The software component to which this billing item points.
	Resource *SoftLayer_Software_Component `json:"resource"`

	// ResourceTableId - The resource (unique identifier) for a software component billing item.
	ResourceTableId int `json:"resourceTableId"`
}
