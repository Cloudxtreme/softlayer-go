package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 - The
// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 data type represents a single
// McAfee Secure anti-virus/spyware software component that uses the ePolicy Orchestrator version 4.5
// backend.
type SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 struct {

	// EpoVersion - The version of ePolicy Orchestrator that the anti-virus/spyware client communicates
	// with.
	EpoVersion string `json:"epoVersion"`
}

func (softlayer_software_component_antivirusspyware_mcafee_epo_version45 *SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45) String() string {
	return "SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45"
}

// SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45_Extended is SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45 with all maskable types.
type SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45_Extended struct {
	SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45

	// DataFileVersion - no documentation
	DataFileVersion *McAfee_Epolicy_Orchestrator_Version45_Product_Properties `json:"dataFileVersion"`

	// LatestSpywareEvents - no documentation
	LatestSpywareEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"latestSpywareEvents"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus"`

	// LatestAccessProtectionEventCount - no documentation
	LatestAccessProtectionEventCount uint64 `json:"latestAccessProtectionEventCount"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails"`

	// CurrentAntivirusPolicy - no documentation
	CurrentAntivirusPolicy int `json:"currentAntivirusPolicy"`

	// LatestAccessProtectionEvents - no documentation
	LatestAccessProtectionEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"latestAccessProtectionEvents"`

	// LatestAntivirusEvents - no documentation
	LatestAntivirusEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"latestAntivirusEvents"`

	// LatestAntivirusEventCount - no documentation
	LatestAntivirusEventCount uint64 `json:"latestAntivirusEventCount"`

	// LatestSpywareEventCount - no documentation
	LatestSpywareEventCount uint64 `json:"latestSpywareEventCount"`
}

func (softlayer_software_component_antivirusspyware_mcafee_epo_version45 *SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45_Extended) String() string {
	return "SoftLayer_Software_Component_AntivirusSpyware_Mcafee_Epo_Version45"
}
