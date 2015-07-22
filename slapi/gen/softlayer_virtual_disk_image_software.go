package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Disk_Image_Software - A SoftLayer_Virtual_Disk_Image_Software record connects a
// computing instance's virtual disk images with software records. This can be useful if a disk image
// is directly associated with software such as operating systems.
type SoftLayer_Virtual_Disk_Image_Software struct {

	// DiskImage - The virtual disk image that is associated with software.
	DiskImage *SoftLayer_Virtual_Disk_Image `json:"diskImage"`

	// Id - The unique identifier of a virtual disk image to software relationship.
	Id int `json:"id"`

	// PasswordCount - A count of username/Password pairs used for access to a Software Installation.
	PasswordCount uint64 `json:"passwordCount"`

	// Passwords - Username/Password pairs used for access to a Software Installation.
	Passwords []*SoftLayer_Virtual_Disk_Image_Software_Password `json:"passwords"`

	// SoftwareDescription - no documentation
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// SoftwareDescriptionId - The unique identifier of the software that a virtual disk image is
	// associated with.
	SoftwareDescriptionId int `json:"softwareDescriptionId"`
}
