package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Disk_Image_Software_Password - This SoftLayer_Virtual_Disk_Image_Software_Password
// data type contains a password for a specific virtual disk image software instance.
type SoftLayer_Virtual_Disk_Image_Software_Password struct {

	// Password - no documentation
	Password string `json:"password"`

	// Software - The instance that this username/password pair is valid for.
	Software *SoftLayer_Virtual_Disk_Image_Software `json:"software"`

	// Username - no documentation
	Username string `json:"username"`
}
