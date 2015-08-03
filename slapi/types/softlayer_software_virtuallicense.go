package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_VirtualLicense - SoftLayer_Software_VirtualLicense is the application class that
// handles a special type of Software License. Most software licenses are licensed to a specific
// hardware ID; virtual licenses are designed for virtual machines and therefore are assigned to an IP
// Address. Not all software packages can be "virtual licensed".
type SoftLayer_Software_VirtualLicense struct {

	// IpAddress - The specific IP address this Virtual License belongs to.
	IpAddress string `json:"ipAddress"`

	// SoftwareDescriptionId - The Software Description ID this Virtual License is for.
	SoftwareDescriptionId int `json:"softwareDescriptionId"`

	// AccountId - The ID of the SoftLayer Account to which this Virtual License belongs to.
	AccountId int `json:"accountId"`

	// Key - no documentation
	Key string `json:"key"`

	// Notes - A "notes" string attached to this specific Virtual License.
	Notes string `json:"notes"`

	// SubnetId - The ID of the SoftLayer Network Subnet this Virtual License belongs to.
	SubnetId int `json:"subnetId"`

	// HostHardwareId - The ID of the SoftLayer Hardware Server record to which this Virtual License
	// belongs.
	HostHardwareId int `json:"hostHardwareId"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_software_virtuallicense *SoftLayer_Software_VirtualLicense) String() string {
	return "SoftLayer_Software_VirtualLicense"
}

// SoftLayer_Software_VirtualLicense_Extended is SoftLayer_Software_VirtualLicense with all maskable types.
type SoftLayer_Software_VirtualLicense_Extended struct {
	SoftLayer_Software_VirtualLicense

	// HostHardware - The hardware record to which the software virtual license is assigned.
	HostHardware *SoftLayer_Hardware_Server `json:"hostHardware"`

	// Account - The customer account this Virtual License belongs to.
	Account *SoftLayer_Account `json:"account"`

	// IpAddressRecord - The IP Address record associated with a virtual license.
	IpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"ipAddressRecord"`

	// Subnet - The subnet this Virtual License's IP address belongs to.
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// SoftwareDescription - The SoftLayer_Software_Description that this virtual license is for.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`
}

func (softlayer_software_virtuallicense *SoftLayer_Software_VirtualLicense_Extended) String() string {
	return "SoftLayer_Software_VirtualLicense"
}
