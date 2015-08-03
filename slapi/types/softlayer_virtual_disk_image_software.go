package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Disk_Image_Software - A SoftLayer_Virtual_Disk_Image_Software record connects a
// computing instance's virtual disk images with software records. This can be useful if a disk image
// is directly associated with software such as operating systems.
type SoftLayer_Virtual_Disk_Image_Software struct {

	// Id - The unique identifier of a virtual disk image to software relationship.
	Id int `json:"id"`

	// SoftwareDescriptionId - The unique identifier of the software that a virtual disk image is
	// associated with.
	SoftwareDescriptionId int `json:"softwareDescriptionId"`
}

// SoftLayer_Virtual_Disk_Image_Software_Extended is SoftLayer_Virtual_Disk_Image_Software with all maskable types.
type SoftLayer_Virtual_Disk_Image_Software_Extended struct {
	SoftLayer_Virtual_Disk_Image_Software

	// SoftwareDescription - no documentation
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// PasswordCount - A count of username/Password pairs used for access to a Software Installation.
	PasswordCount uint64 `json:"passwordCount"`

	// DiskImage - The virtual disk image that is associated with software.
	DiskImage *SoftLayer_Virtual_Disk_Image `json:"diskImage"`

	// Passwords - Username/Password pairs used for access to a Software Installation.
	Passwords []*SoftLayer_Virtual_Disk_Image_Software_Password `json:"passwords"`
}

func (softlayer_virtual_disk_image_software *SoftLayer_Virtual_Disk_Image_Software) String() string {
	return "SoftLayer_Virtual_Disk_Image_Software"
}
