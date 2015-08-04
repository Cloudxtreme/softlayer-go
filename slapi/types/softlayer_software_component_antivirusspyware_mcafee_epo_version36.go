package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 - The
// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 data type represents a single
// McAfee Secure anti-virus/spyware software component that uses the ePolicy Orchestrator version 3.6
// backend.
type SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 struct {

	// HardwareId - Hardware Identification Number for the server this Software Component is installed
	// upon.
	HardwareId int `json:"hardwareId,omitempty"`

	// ManufacturerLicenseInstance - A license key for this specific installation of software, if it is
	// needed.
	ManufacturerLicenseInstance string `json:"manufacturerLicenseInstance,omitempty"`

	// ManufacturerActivationCode - The manufacturer code that is needed to activate a license.
	ManufacturerActivationCode string `json:"manufacturerActivationCode,omitempty"`

	// Id - An ID number identifying this Software Component (Software Installation)
	Id int `json:"id,omitempty"`

	// EpoVersion - The version of ePolicy Orchestrator that the anti-virus/spyware client communicates
	// with.
	EpoVersion string `json:"epoVersion,omitempty"`

	// SoftwareDescription - The Software Description of this Software Component.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// CurrentAntivirusPolicy - no documentation
	CurrentAntivirusPolicy int `json:"currentAntivirusPolicy,omitempty"`

	// LatestAccessProtectionEvents - no documentation
	LatestAccessProtectionEvents []*McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection `json:"latestAccessProtectionEvents,omitempty"`

	// LatestAntivirusEvents - no documentation
	LatestAntivirusEvents []*McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event `json:"latestAntivirusEvents,omitempty"`

	// PasswordHistoryCount - no documentation
	PasswordHistoryCount uint64 `json:"passwordHistoryCount,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// LatestAccessProtectionEventCount - no documentation
	LatestAccessProtectionEventCount uint64 `json:"latestAccessProtectionEventCount,omitempty"`

	// Hardware - The hardware this Software Component is installed upon.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// PasswordCount - A count of username/Password pairs used for access to this Software Installation.
	PasswordCount uint64 `json:"passwordCount,omitempty"`

	// PasswordHistory - no documentation
	PasswordHistory []*SoftLayer_Software_Component_Password_History `json:"passwordHistory,omitempty"`

	// VirtualGuest - The virtual guest this software component is installed upon.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// LatestSpywareEventCount - no documentation
	LatestSpywareEventCount uint64 `json:"latestSpywareEventCount,omitempty"`

	// SoftwareLicense - no documentation
	SoftwareLicense *SoftLayer_Software_License `json:"softwareLicense,omitempty"`

	// AverageInstallationDuration - The average amount of time that a software component takes to install.
	AverageInstallationDuration uint64 `json:"averageInstallationDuration,omitempty"`

	// Passwords - Username/Password pairs used for access to this Software Installation.
	Passwords []*SoftLayer_Software_Component_Password `json:"passwords,omitempty"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version36_Agent_Details `json:"agentDetails,omitempty"`

	// DataFileVersion - no documentation
	DataFileVersion *McAfee_Epolicy_Orchestrator_Version36_Product_Properties `json:"dataFileVersion,omitempty"`

	// LatestSpywareEvents - no documentation
	LatestSpywareEvents []*McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event `json:"latestSpywareEvents,omitempty"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus,omitempty"`

	// LatestAntivirusEventCount - no documentation
	LatestAntivirusEventCount uint64 `json:"latestAntivirusEventCount,omitempty"`
}

func (softlayer_software_component_antivirusspyware_mcafee_epo_version36 *SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) String() string {
	return "SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36"
}
