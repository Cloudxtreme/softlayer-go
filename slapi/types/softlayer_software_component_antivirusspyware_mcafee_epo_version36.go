package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 - The
// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 data type represents a single
// McAfee Secure anti-virus/spyware software component that uses the ePolicy Orchestrator version 3.6
// backend.
type SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 struct {

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version36_Agent_Details `json:"agentDetails"`

	// CurrentAntivirusPolicy - no documentation
	CurrentAntivirusPolicy int `json:"currentAntivirusPolicy"`

	// DataFileVersion - no documentation
	DataFileVersion *McAfee_Epolicy_Orchestrator_Version36_Product_Properties `json:"dataFileVersion"`

	// EpoVersion - The version of ePolicy Orchestrator that the anti-virus/spyware client communicates
	// with.
	EpoVersion string `json:"epoVersion"`

	// LatestAccessProtectionEventCount - no documentation
	LatestAccessProtectionEventCount uint64 `json:"latestAccessProtectionEventCount"`

	// LatestAccessProtectionEvents - no documentation
	LatestAccessProtectionEvents []*McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection `json:"latestAccessProtectionEvents"`

	// LatestAntivirusEventCount - no documentation
	LatestAntivirusEventCount uint64 `json:"latestAntivirusEventCount"`

	// LatestAntivirusEvents - no documentation
	LatestAntivirusEvents []*McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event `json:"latestAntivirusEvents"`

	// LatestSpywareEventCount - no documentation
	LatestSpywareEventCount uint64 `json:"latestSpywareEventCount"`

	// LatestSpywareEvents - no documentation
	LatestSpywareEvents []*McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event `json:"latestSpywareEvents"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus"`
}

func (softlayer_software_component_antivirusspyware_mcafee_epo_version36 *SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) String() string {
	return "SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36"
}
