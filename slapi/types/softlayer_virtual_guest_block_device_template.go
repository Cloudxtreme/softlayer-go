package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Guest_Block_Device_Template - The virtual block device template data type presents
// the structure in which all archived image templates are presented. A virtual block device template,
// also known as a image template, represents the image of a virtual guest instance.
type SoftLayer_Virtual_Guest_Block_Device_Template struct {

	// Device - no documentation
	Device string `json:"device"`

	// DiskImage - no documentation
	DiskImage *SoftLayer_Virtual_Disk_Image `json:"diskImage"`

	// DiskImageId - A block device template's [[SoftLayer_Virtual_Disk_Image|disk image]] ID.
	DiskImageId int `json:"diskImageId"`

	// DiskSpace - The amount of disk space that a block device template is using. Use this number along
	// with the units property to obtain the correct space used.
	DiskSpace float32 `json:"diskSpace"`

	// Group - A block device template's group. Several block device templates can be combined together
	// into a group for archiving purposes.
	Group *SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"group"`

	// GroupId - A block device template's [[SoftLayer_Virtual_Guest_Block_Device_Template_Group|group]]
	// ID.
	GroupId int `json:"groupId"`

	// Id - no documentation
	Id int `json:"id"`

	// Units - The units that will be used with the disk space property to identify the amount of disk
	// space used.
	Units string `json:"units"`
}

func (softlayer_virtual_guest_block_device_template *SoftLayer_Virtual_Guest_Block_Device_Template) String() string {
	return "SoftLayer_Virtual_Guest_Block_Device_Template"
}
