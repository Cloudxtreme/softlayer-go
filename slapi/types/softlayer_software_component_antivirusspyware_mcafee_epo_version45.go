package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 - The
// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 data type represents a single
// McAfee Secure anti-virus/spyware software component that uses the ePolicy Orchestrator version 4.5
// backend.
type SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 struct {

	// EpoVersion - The version of ePolicy Orchestrator that the anti-virus/spyware client communicates
	// with.
	EpoVersion string `json:"epoVersion,omitempty"`

	// ManufacturerActivationCode - The manufacturer code that is needed to activate a license.
	ManufacturerActivationCode string `json:"manufacturerActivationCode,omitempty"`

	// ManufacturerLicenseInstance - A license key for this specific installation of software, if it is
	// needed.
	ManufacturerLicenseInstance string `json:"manufacturerLicenseInstance,omitempty"`

	// Id - An ID number identifying this Software Component (Software Installation)
	Id int `json:"id,omitempty"`

	// HardwareId - Hardware Identification Number for the server this Software Component is installed
	// upon.
	HardwareId int `json:"hardwareId,omitempty"`
}

func (softlayer_software_component_antivirusspyware_mcafee_epo_version45 *SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45) String() string {
	return "SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45"
}

// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45_Extended is SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 with all maskable types.
type SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45_Extended struct {
	SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45

	// DataFileVersion - no documentation
	DataFileVersion *McAfee_Epolicy_Orchestrator_Version45_Product_Properties `json:"dataFileVersion,omitempty"`

	// LatestAntivirusEvents - no documentation
	LatestAntivirusEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"latestAntivirusEvents,omitempty"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus,omitempty"`

	// PasswordCount - A count of username/Password pairs used for access to this Software Installation.
	PasswordCount uint64 `json:"passwordCount,omitempty"`

	// LatestSpywareEvents - no documentation
	LatestSpywareEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"latestSpywareEvents,omitempty"`

	// LatestAccessProtectionEventCount - no documentation
	LatestAccessProtectionEventCount uint64 `json:"latestAccessProtectionEventCount,omitempty"`

	// LatestSpywareEventCount - no documentation
	LatestSpywareEventCount uint64 `json:"latestSpywareEventCount,omitempty"`

	// PasswordHistoryCount - no documentation
	PasswordHistoryCount uint64 `json:"passwordHistoryCount,omitempty"`

	// LatestAccessProtectionEvents - no documentation
	LatestAccessProtectionEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"latestAccessProtectionEvents,omitempty"`

	// LatestAntivirusEventCount - no documentation
	LatestAntivirusEventCount uint64 `json:"latestAntivirusEventCount,omitempty"`

	// Passwords - Username/Password pairs used for access to this Software Installation.
	Passwords []*SoftLayer_Software_Component_Password `json:"passwords,omitempty"`

	// SoftwareDescription - The Software Description of this Software Component.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// SoftwareLicense - no documentation
	SoftwareLicense *SoftLayer_Software_License `json:"softwareLicense,omitempty"`

	// PasswordHistory - no documentation
	PasswordHistory []*SoftLayer_Software_Component_Password_History `json:"passwordHistory,omitempty"`

	// AverageInstallationDuration - The average amount of time that a software component takes to install.
	AverageInstallationDuration uint64 `json:"averageInstallationDuration,omitempty"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails,omitempty"`

	// CurrentAntivirusPolicy - no documentation
	CurrentAntivirusPolicy int `json:"currentAntivirusPolicy,omitempty"`

	// Hardware - The hardware this Software Component is installed upon.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// VirtualGuest - The virtual guest this software component is installed upon.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`
}

func (softlayer_software_component_antivirusspyware_mcafee_epo_version45 *SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45_Extended) String() string {
	return "SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45"
}
