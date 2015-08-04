package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Disk_Image_Type - SoftLayer_Virtual_Disk_Image_Type models the types of virtual
// disk images available to CloudLayer Computing Instances. Virtual disk image types describe if an
// image's data is preservable when upgraded, whether a disk contains a suspended virtual image, or if
// a disk contains crash dump information.
type SoftLayer_Virtual_Disk_Image_Type struct {

	// Description - A brief description of a virtual disk image type's function.
	Description string `json:"description,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_virtual_disk_image_type *SoftLayer_Virtual_Disk_Image_Type) String() string {
	return "SoftLayer_Virtual_Disk_Image_Type"
}
