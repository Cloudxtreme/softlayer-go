package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 - The
// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 data type represents a single
// McAfee Secure anti-virus/spyware software component that uses the ePolicy Orchestrator version 3.6
// backend.
type SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 struct {

	// EpoVersion - The version of ePolicy Orchestrator that the anti-virus/spyware client communicates
	// with.
	EpoVersion string `json:"epoVersion"`
}

// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36_Extended is SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36 with all maskable types.
type SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36_Extended struct {
	SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36

	// LatestAccessProtectionEventCount - no documentation
	LatestAccessProtectionEventCount uint64 `json:"latestAccessProtectionEventCount"`

	// LatestAntivirusEventCount - no documentation
	LatestAntivirusEventCount uint64 `json:"latestAntivirusEventCount"`

	// CurrentAntivirusPolicy - no documentation
	CurrentAntivirusPolicy int `json:"currentAntivirusPolicy"`

	// LatestAccessProtectionEvents - no documentation
	LatestAccessProtectionEvents []*McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event_AccessProtection `json:"latestAccessProtectionEvents"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus"`

	// LatestSpywareEventCount - no documentation
	LatestSpywareEventCount uint64 `json:"latestSpywareEventCount"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version36_Agent_Details `json:"agentDetails"`

	// DataFileVersion - no documentation
	DataFileVersion *McAfee_Epolicy_Orchestrator_Version36_Product_Properties `json:"dataFileVersion"`

	// LatestAntivirusEvents - no documentation
	LatestAntivirusEvents []*McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event `json:"latestAntivirusEvents"`

	// LatestSpywareEvents - no documentation
	LatestSpywareEvents []*McAfee_Epolicy_Orchestrator_Version36_Antivirus_Event `json:"latestSpywareEvents"`
}

func (softlayer_software_component_antivirusspyware_mcafee_epo_version36 *SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36) String() string {
	return "SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version36"
}
