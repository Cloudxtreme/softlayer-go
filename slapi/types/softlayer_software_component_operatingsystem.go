package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Software_Component_OperatingSystem - SoftLayer_Software_Component_OperatingSystem extends
// the [[SoftLayer_Software_Component]] data type to include operating system specific properties.
type SoftLayer_Software_Component_OperatingSystem struct {

	// Id - An ID number identifying this Software Component (Software Installation)
	Id int `json:"id,omitempty"`

	// ManufacturerActivationCode - The manufacturer code that is needed to activate a license.
	ManufacturerActivationCode string `json:"manufacturerActivationCode,omitempty"`

	// ManufacturerLicenseInstance - A license key for this specific installation of software, if it is
	// needed.
	ManufacturerLicenseInstance string `json:"manufacturerLicenseInstance,omitempty"`

	// HardwareId - Hardware Identification Number for the server this Software Component is installed
	// upon.
	HardwareId int `json:"hardwareId,omitempty"`

	// ReloadTransactionGroup - An operating systems associated
	// [[SoftLayer_Provisioning_Version1_Transaction_Group|Transaction Group]]. A transaction group is a
	// list of operations that will occur during the installment of an operating system.
	ReloadTransactionGroup *SoftLayer_Provisioning_Version1_Transaction_Group `json:"reloadTransactionGroup,omitempty"`

	// PasswordHistory - no documentation
	PasswordHistory []*SoftLayer_Software_Component_Password_History `json:"passwordHistory,omitempty"`

	// PartitionTemplateCount - A count of an operating system's associated
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]] that can be used to
	// configure a hardware drive.
	PartitionTemplateCount uint64 `json:"partitionTemplateCount,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// Hardware - The hardware this Software Component is installed upon.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// Passwords - Username/Password pairs used for access to this Software Installation.
	Passwords []*SoftLayer_Software_Component_Password `json:"passwords,omitempty"`

	// PasswordHistoryCount - no documentation
	PasswordHistoryCount uint64 `json:"passwordHistoryCount,omitempty"`

	// LicenseExpirationDate - The date in which the license for this software expires.
	LicenseExpirationDate *time.Time `json:"licenseExpirationDate,omitempty"`

	// PartitionTemplates - An operating system's associated
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]] that can be used to
	// configure a hardware drive.
	PartitionTemplates []*SoftLayer_Hardware_Component_Partition_Template `json:"partitionTemplates,omitempty"`

	// VirtualGuest - The virtual guest this software component is installed upon.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// SoftwareLicense - no documentation
	SoftwareLicense *SoftLayer_Software_License `json:"softwareLicense,omitempty"`

	// PasswordCount - A count of username/Password pairs used for access to this Software Installation.
	PasswordCount uint64 `json:"passwordCount,omitempty"`

	// SoftwareDescription - The Software Description of this Software Component.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// AverageInstallationDuration - The average amount of time that a software component takes to install.
	AverageInstallationDuration uint64 `json:"averageInstallationDuration,omitempty"`
}

func (softlayer_software_component_operatingsystem *SoftLayer_Software_Component_OperatingSystem) String() string {
	return "SoftLayer_Software_Component_OperatingSystem"
}
