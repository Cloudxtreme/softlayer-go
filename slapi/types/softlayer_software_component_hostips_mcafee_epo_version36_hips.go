package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips - The
// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips data type represents a single McAfee
// Secure Host IPS software component that uses the ePolicy Orchestrator version 3.6 backend.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips struct {

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version36_Agent_Details `json:"agentDetails"`

	// ApplicationModePolicyNameCount - A count of the names of the possible policy options for the
	// application mode setting.
	ApplicationModePolicyNameCount uint64 `json:"applicationModePolicyNameCount"`

	// ApplicationModePolicyNames - The names of the possible policy options for the application mode
	// setting.
	ApplicationModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationModePolicyNames"`

	// ApplicationRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// application rule set setting.
	ApplicationRuleSetPolicyNameCount uint64 `json:"applicationRuleSetPolicyNameCount"`

	// ApplicationRuleSetPolicyNames - The names of the possible policy options for the application rule
	// set setting.
	ApplicationRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationRuleSetPolicyNames"`

	// EnforcementPolicyNameCount - A count of the names of the possible options for the enforcement policy
	// setting.
	EnforcementPolicyNameCount uint64 `json:"enforcementPolicyNameCount"`

	// EnforcementPolicyNames - The names of the possible options for the enforcement policy setting.
	EnforcementPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"enforcementPolicyNames"`

	// EpoVersion - The version of ePolicy Orchestrator that the host IPS client communicates with.
	EpoVersion string `json:"epoVersion"`

	// FirewallModePolicyNameCount - A count of the names of the possible policy options for the firewall
	// mode setting.
	FirewallModePolicyNameCount uint64 `json:"firewallModePolicyNameCount"`

	// FirewallModePolicyNames - The names of the possible policy options for the firewall mode setting.
	FirewallModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallModePolicyNames"`

	// FirewallRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// firewall rule set setting.
	FirewallRuleSetPolicyNameCount uint64 `json:"firewallRuleSetPolicyNameCount"`

	// FirewallRuleSetPolicyNames - The names of the possible policy options for the firewall rule set
	// setting.
	FirewallRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallRuleSetPolicyNames"`

	// IpsModePolicyNameCount - A count of the names of the possible policy options for the host IPS mode
	// setting.
	IpsModePolicyNameCount uint64 `json:"ipsModePolicyNameCount"`

	// IpsModePolicyNames - The names of the possible policy options for the host IPS mode setting.
	IpsModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsModePolicyNames"`

	// IpsProtectionPolicyNameCount - A count of the names of the possible policy options for the host IPS
	// protection setting.
	IpsProtectionPolicyNameCount uint64 `json:"ipsProtectionPolicyNameCount"`

	// IpsProtectionPolicyNames - The names of the possible policy options for the host IPS protection
	// setting.
	IpsProtectionPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsProtectionPolicyNames"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus"`
}

func (softlayer_software_component_hostips_mcafee_epo_version36_hips *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips"
}
