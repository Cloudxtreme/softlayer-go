package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat - The
// SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Microsoft data type contains
// general information relating to a single SoftLayer billing item for a Microsoft operating system
// software components on virtual machines.
type SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat struct {

	// ResourceTableId - The resource (unique identifier) for a software component billing item.
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_billing_item_software_component_virtual_operatingsystem_redhat *SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat) String() string {
	return "SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat"
}

// SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat_Extended is SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat with all maskable types.
type SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat_Extended struct {
	SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat

	// Resource - The software component to which this billing item points.
	Resource *SoftLayer_Software_Component `json:"resource"`
}

func (softlayer_billing_item_software_component_virtual_operatingsystem_redhat *SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat_Extended) String() string {
	return "SoftLayer_Billing_Item_Software_Component_Virtual_OperatingSystem_Redhat"
}
