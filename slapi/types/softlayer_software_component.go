package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component - A SoftLayer_Software_Component ties the installation of a specific
// piece of software onto a specific piece of hardware. SoftLayer_Software_Component works with
// SoftLayer_Software_License and SoftLayer_Software_Description to tie this all together.
// SoftLayer_Software_Component is the installation of a specific piece of software onto a specific
// piece of hardware in accordance to a software license. SoftLayer_Software_License dictates when and
// how a specific piece of software may be installed onto a piece of hardware.
// SoftLayer_Software_Description describes a specific piece of software which can be installed onto
// hardware in accordance with it's license agreement.
type SoftLayer_Software_Component struct {

	// ManufacturerActivationCode - The manufacturer code that is needed to activate a license.
	ManufacturerActivationCode string `json:"manufacturerActivationCode,omitempty"`

	// HardwareId - Hardware Identification Number for the server this Software Component is installed
	// upon.
	HardwareId int `json:"hardwareId,omitempty"`

	// ManufacturerLicenseInstance - A license key for this specific installation of software, if it is
	// needed.
	ManufacturerLicenseInstance string `json:"manufacturerLicenseInstance,omitempty"`

	// Id - An ID number identifying this Software Component (Software Installation)
	Id int `json:"id,omitempty"`
}

func (softlayer_software_component *SoftLayer_Software_Component) String() string {
	return "SoftLayer_Software_Component"
}

// SoftLayer_Software_Component_Extended is SoftLayer_Software_Component with all maskable types.
type SoftLayer_Software_Component_Extended struct {
	SoftLayer_Software_Component

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// SoftwareDescription - The Software Description of this Software Component.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// AverageInstallationDuration - The average amount of time that a software component takes to install.
	AverageInstallationDuration uint64 `json:"averageInstallationDuration,omitempty"`

	// Hardware - The hardware this Software Component is installed upon.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// SoftwareLicense - no documentation
	SoftwareLicense *SoftLayer_Software_License `json:"softwareLicense,omitempty"`

	// PasswordCount - A count of username/Password pairs used for access to this Software Installation.
	PasswordCount uint64 `json:"passwordCount,omitempty"`

	// Passwords - Username/Password pairs used for access to this Software Installation.
	Passwords []*SoftLayer_Software_Component_Password `json:"passwords,omitempty"`

	// VirtualGuest - The virtual guest this software component is installed upon.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// PasswordHistoryCount - no documentation
	PasswordHistoryCount uint64 `json:"passwordHistoryCount,omitempty"`

	// PasswordHistory - no documentation
	PasswordHistory []*SoftLayer_Software_Component_Password_History `json:"passwordHistory,omitempty"`
}

func (softlayer_software_component *SoftLayer_Software_Component_Extended) String() string {
	return "SoftLayer_Software_Component"
}
