package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_VirtualLicense - SoftLayer_Software_VirtualLicense is the application class that
// handles a special type of Software License. Most software licenses are licensed to a specific
// hardware ID; virtual licenses are designed for virtual machines and therefore are assigned to an IP
// Address. Not all software packages can be "virtual licensed".
type SoftLayer_Software_VirtualLicense struct {

	// Account - The customer account this Virtual License belongs to.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The ID of the SoftLayer Account to which this Virtual License belongs to.
	AccountId int `json:"accountId"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// HostHardware - The hardware record to which the software virtual license is assigned.
	HostHardware *SoftLayer_Hardware_Server `json:"hostHardware"`

	// HostHardwareId - The ID of the SoftLayer Hardware Server record to which this Virtual License
	// belongs.
	HostHardwareId int `json:"hostHardwareId"`

	// Id - no documentation
	Id int `json:"id"`

	// IpAddress - The specific IP address this Virtual License belongs to.
	IpAddress string `json:"ipAddress"`

	// IpAddressRecord - The IP Address record associated with a virtual license.
	IpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"ipAddressRecord"`

	// Key - no documentation
	Key string `json:"key"`

	// Notes - A "notes" string attached to this specific Virtual License.
	Notes string `json:"notes"`

	// SoftwareDescription - The SoftLayer_Software_Description that this virtual license is for.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// SoftwareDescriptionId - The Software Description ID this Virtual License is for.
	SoftwareDescriptionId int `json:"softwareDescriptionId"`

	// Subnet - The subnet this Virtual License's IP address belongs to.
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`

	// SubnetId - The ID of the SoftLayer Network Subnet this Virtual License belongs to.
	SubnetId int `json:"subnetId"`
}

// GetLicenseFile - Attempt to retrieve the file associated with a virtual license, if such a file
// exists. If there is no file for this virtual license, calling this method will either throw an
// exception or return false.
func (softlayer_software_virtuallicense *SoftLayer_Software_VirtualLicense) GetLicenseFile() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Software_VirtualLicense object whose ID number
// corresponds to the ID number of the init parameter passed to the SoftLayer_Software_VirtualLicense
// service. You can only retrieve Virtual Licenses assigned to your account number.
func (softlayer_software_virtuallicense *SoftLayer_Software_VirtualLicense) GetObject() (*SoftLayer_Software_VirtualLicense, error) {
	var returnValue *SoftLayer_Software_VirtualLicense
	return returnValue, nil
}
