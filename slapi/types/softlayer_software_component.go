package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Software_Component - A SoftLayer_Software_Component ties the installation of a specific
// piece of software onto a specific piece of hardware. SoftLayer_Software_Component works with
// SoftLayer_Software_License and SoftLayer_Software_Description to tie this all together.
// SoftLayer_Software_Component is the installation of a specific piece of software onto a specific
// piece of hardware in accordance to a software license. SoftLayer_Software_License dictates when and
// how a specific piece of software may be installed onto a piece of hardware.
// SoftLayer_Software_Description describes a specific piece of software which can be installed onto
// hardware in accordance with it's license agreement.
type SoftLayer_Software_Component struct {

	// AverageInstallationDuration - The average amount of time that a software component takes to install.
	AverageInstallationDuration uint64 `json:"averageInstallationDuration"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// Hardware - The hardware this Software Component is installed upon.
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - Hardware Identification Number for the server this Software Component is installed
	// upon.
	HardwareId int `json:"hardwareId"`

	// Id - An ID number identifying this Software Component (Software Installation)
	Id int `json:"id"`

	// ManufacturerActivationCode - The manufacturer code that is needed to activate a license.
	ManufacturerActivationCode string `json:"manufacturerActivationCode"`

	// ManufacturerLicenseInstance - A license key for this specific installation of software, if it is
	// needed.
	ManufacturerLicenseInstance string `json:"manufacturerLicenseInstance"`

	// PasswordCount - A count of username/Password pairs used for access to this Software Installation.
	PasswordCount uint64 `json:"passwordCount"`

	// PasswordHistory - no documentation
	PasswordHistory []*SoftLayer_Software_Component_Password_History `json:"passwordHistory"`

	// PasswordHistoryCount - no documentation
	PasswordHistoryCount uint64 `json:"passwordHistoryCount"`

	// Passwords - Username/Password pairs used for access to this Software Installation.
	Passwords []*SoftLayer_Software_Component_Password `json:"passwords"`

	// SoftwareDescription - The Software Description of this Software Component.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription"`

	// SoftwareLicense - no documentation
	SoftwareLicense *SoftLayer_Software_License `json:"softwareLicense"`

	// VirtualGuest - The virtual guest this software component is installed upon.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`
}

// GetLicenseFile - Attempt to retrieve the file associated with a software component. If the software
// component does not support downloading license files an exception will be thrown.
func (softlayer_software_component *SoftLayer_Software_Component) GetLicenseFile(commonOptions *slapi.CommonOptions) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Software_Component object whose ID corresponds to the
// ID number of the init parameter passed to the SoftLayer_Software_Component service. The best way to
// get software components is through getSoftwareComponents from the Hardware service.
func (softlayer_software_component *SoftLayer_Software_Component) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Software_Component, error) {
	var returnValue *SoftLayer_Software_Component
	return returnValue, nil
}

// GetVendorSetUpConfiguration - <nil>
func (softlayer_software_component *SoftLayer_Software_Component) GetVendorSetUpConfiguration(commonOptions *slapi.CommonOptions) (string, error) {
	var returnValue string
	return returnValue, nil
}