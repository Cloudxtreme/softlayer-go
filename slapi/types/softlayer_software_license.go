package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_License - This class describes a specific type of license, like a Microsoft
// Windows Site License, a GPL license, or a license of another type.
type SoftLayer_Software_License struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// SoftwareDescriptionId - The ID number of a Software Description that this specific license is valid
	// for.
	SoftwareDescriptionId int `json:"softwareDescriptionId,omitempty"`
}

func (softlayer_software_license *SoftLayer_Software_License) String() string {
	return "SoftLayer_Software_License"
}

// SoftLayer_Software_License_Extended is SoftLayer_Software_License with all maskable types.
type SoftLayer_Software_License_Extended struct {
	SoftLayer_Software_License

	// Account - The account that owns this specific License instance.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Owner - The account that owns this specific License instance.
	Owner *SoftLayer_Account `json:"owner,omitempty"`

	// SoftwareDescription - A Description of the software that this license instance is valid for.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`
}

func (softlayer_software_license *SoftLayer_Software_License_Extended) String() string {
	return "SoftLayer_Software_License"
}
