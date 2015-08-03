package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group - <nil>
type SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group struct {
}

// SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group_Extended is SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group with all maskable types.
type SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group_Extended struct {
	SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group

	// Resource - <nil>
	Resource *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"resource"`
}

func (softlayer_tag_reference_virtual_guest_block_device_template_group *SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group) String() string {
	return "SoftLayer_Tag_Reference_Virtual_Guest_Block_Device_Template_Group"
}
