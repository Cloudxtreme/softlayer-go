package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips - The
// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips data type represents a single McAfee
// Secure Host IPS software component that uses the ePolicy Orchestrator version 3.6 backend.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips struct {

	// EpoVersion - The version of ePolicy Orchestrator that the host IPS client communicates with.
	EpoVersion string `json:"epoVersion"`
}

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Extended is SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips with all maskable types.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Extended struct {
	SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips

	// EnforcementPolicyNameCount - A count of the names of the possible options for the enforcement policy
	// setting.
	EnforcementPolicyNameCount uint64 `json:"enforcementPolicyNameCount"`

	// IpsModePolicyNameCount - A count of the names of the possible policy options for the host IPS mode
	// setting.
	IpsModePolicyNameCount uint64 `json:"ipsModePolicyNameCount"`

	// IpsProtectionPolicyNameCount - A count of the names of the possible policy options for the host IPS
	// protection setting.
	IpsProtectionPolicyNameCount uint64 `json:"ipsProtectionPolicyNameCount"`

	// ApplicationRuleSetPolicyNames - The names of the possible policy options for the application rule
	// set setting.
	ApplicationRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationRuleSetPolicyNames"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus"`

	// IpsModePolicyNames - The names of the possible policy options for the host IPS mode setting.
	IpsModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsModePolicyNames"`

	// ApplicationModePolicyNameCount - A count of the names of the possible policy options for the
	// application mode setting.
	ApplicationModePolicyNameCount uint64 `json:"applicationModePolicyNameCount"`

	// IpsProtectionPolicyNames - The names of the possible policy options for the host IPS protection
	// setting.
	IpsProtectionPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsProtectionPolicyNames"`

	// EnforcementPolicyNames - The names of the possible options for the enforcement policy setting.
	EnforcementPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"enforcementPolicyNames"`

	// FirewallModePolicyNames - The names of the possible policy options for the firewall mode setting.
	FirewallModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallModePolicyNames"`

	// FirewallRuleSetPolicyNames - The names of the possible policy options for the firewall rule set
	// setting.
	FirewallRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallRuleSetPolicyNames"`

	// ApplicationRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// application rule set setting.
	ApplicationRuleSetPolicyNameCount uint64 `json:"applicationRuleSetPolicyNameCount"`

	// FirewallModePolicyNameCount - A count of the names of the possible policy options for the firewall
	// mode setting.
	FirewallModePolicyNameCount uint64 `json:"firewallModePolicyNameCount"`

	// FirewallRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// firewall rule set setting.
	FirewallRuleSetPolicyNameCount uint64 `json:"firewallRuleSetPolicyNameCount"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version36_Agent_Details `json:"agentDetails"`

	// ApplicationModePolicyNames - The names of the possible policy options for the application mode
	// setting.
	ApplicationModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationModePolicyNames"`
}

func (softlayer_software_component_hostips_mcafee_epo_version36_hips *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips"
}
