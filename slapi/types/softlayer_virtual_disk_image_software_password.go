package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Disk_Image_Software_Password - This SoftLayer_Virtual_Disk_Image_Software_Password
// data type contains a password for a specific virtual disk image software instance.
type SoftLayer_Virtual_Disk_Image_Software_Password struct {

	// Password - no documentation
	Password string `json:"password"`

	// Username - no documentation
	Username string `json:"username"`
}

// SoftLayer_Virtual_Disk_Image_Software_Password_Extended is SoftLayer_Virtual_Disk_Image_Software_Password with all maskable types.
type SoftLayer_Virtual_Disk_Image_Software_Password_Extended struct {
	SoftLayer_Virtual_Disk_Image_Software_Password

	// Software - The instance that this username/password pair is valid for.
	Software *SoftLayer_Virtual_Disk_Image_Software `json:"software"`
}

func (softlayer_virtual_disk_image_software_password *SoftLayer_Virtual_Disk_Image_Software_Password) String() string {
	return "SoftLayer_Virtual_Disk_Image_Software_Password"
}
