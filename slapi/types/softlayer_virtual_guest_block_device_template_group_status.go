package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status - The virtual block device template group
// status data type represents the current status of the image template. Depending upon the status, the
// image template can be used for provisioning or reloading. For an operating system reload, the image
// template will need to have a status of 'Active' or 'Deprecated'. For a provision, the image template
// will need to have a status of 'Active'
type SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status struct {

	// Description - <nil>
	Description string `json:"description"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_virtual_guest_block_device_template_group_status *SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status"
}
