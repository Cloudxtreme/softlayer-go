package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Guest_Block_Device_Template - The virtual block device template data type presents
// the structure in which all archived image templates are presented. A virtual block device template,
// also known as a image template, represents the image of a virtual guest instance.
type SoftLayer_Virtual_Guest_Block_Device_Template struct {

	// GroupId - A block device template's [[SoftLayer_Virtual_Guest_Block_Device_Template_Group|group]]
	// ID.
	GroupId int `json:"groupId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Units - The units that will be used with the disk space property to identify the amount of disk
	// space used.
	Units string `json:"units,omitempty"`

	// Device - no documentation
	Device string `json:"device,omitempty"`

	// DiskImageId - A block device template's [[SoftLayer_Virtual_Disk_Image|disk image]] ID.
	DiskImageId int `json:"diskImageId,omitempty"`

	// DiskSpace - The amount of disk space that a block device template is using. Use this number along
	// with the units property to obtain the correct space used.
	DiskSpace float32 `json:"diskSpace,omitempty"`
}

func (softlayer_virtual_guest_block_device_template *SoftLayer_Virtual_Guest_Block_Device_Template) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template"
}

// SoftLayer_Virtual_Guest_Block_Device_Template_Extended is SoftLayer_Virtual_Guest_Block_Device_Template with all maskable types.
type SoftLayer_Virtual_Guest_Block_Device_Template_Extended struct {
	SoftLayer_Virtual_Guest_Block_Device_Template

	// DiskImage - no documentation
	DiskImage *SoftLayer_Virtual_Disk_Image `json:"diskImage,omitempty"`

	// Group - A block device template's group. Several block device templates can be combined together
	// into a group for archiving purposes.
	Group *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"group,omitempty"`
}

func (softlayer_virtual_guest_block_device_template *SoftLayer_Virtual_Guest_Block_Device_Template_Extended) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template"
}
