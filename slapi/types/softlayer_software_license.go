package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_License - This class describes a specific type of license, like a Microsoft
// Windows Site License, a GPL license, or a license of another type.
type SoftLayer_Software_License struct {

	// Account - The account that owns this specific License instance.
	Account *SoftLayer_Account `json:"account"`

	// Id - no documentation
	Id int `json:"id"`

	// Owner - The account that owns this specific License instance.
	Owner *SoftLayer_Account `json:"owner"`

	// SoftwareDescription - A Description of the software that this license instance is valid for.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// SoftwareDescriptionId - The ID number of a Software Description that this specific license is valid
	// for.
	SoftwareDescriptionId int `json:"softwareDescriptionId"`
}

func (softlayer_software_license *SoftLayer_Software_License) String() string {
	return "SoftLayer_Software_License"
}
