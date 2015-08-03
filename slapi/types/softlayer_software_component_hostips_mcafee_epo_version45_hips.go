package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips - The
// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips data type represents a single McAfee
// Secure Host IPS software component that uses the ePolicy Orchestrator version 4.5 backend.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips struct {

	// EpoVersion - The version of ePolicy Orchestrator that the host IPS client communicates with.
	EpoVersion string `json:"epoVersion"`
}

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Extended is SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips with all maskable types.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Extended struct {
	SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails"`

	// IpsProtectionPolicyNames - The names of the possible policy options for the host IPS protection
	// setting.
	IpsProtectionPolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"ipsProtectionPolicyNames"`

	// EnforcementPolicyNameCount - A count of the names of the possible options for the enforcement policy
	// setting.
	EnforcementPolicyNameCount uint64 `json:"enforcementPolicyNameCount"`

	// IpsEventCount - A count of the host IPS events for this software component.
	IpsEventCount uint64 `json:"ipsEventCount"`

	// EnforcementPolicyNames - The names of the possible options for the enforcement policy setting.
	EnforcementPolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"enforcementPolicyNames"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus"`

	// ApplicationModePolicyNameCount - A count of the names of the possible policy options for the
	// application mode setting.
	ApplicationModePolicyNameCount uint64 `json:"applicationModePolicyNameCount"`

	// BlockedApplicationEventCount - A count of the blocked application events for this software
	// component.
	BlockedApplicationEventCount uint64 `json:"blockedApplicationEventCount"`

	// FirewallRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// firewall rule set setting.
	FirewallRuleSetPolicyNameCount uint64 `json:"firewallRuleSetPolicyNameCount"`

	// IpsModePolicyNameCount - A count of the names of the possible policy options for the host IPS mode
	// setting.
	IpsModePolicyNameCount uint64 `json:"ipsModePolicyNameCount"`

	// ApplicationRuleSetPolicyNames - The names of the possible policy options for the application rule
	// set setting.
	ApplicationRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"applicationRuleSetPolicyNames"`

	// BlockedApplicationEvents - The blocked application events for this software component.
	BlockedApplicationEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"blockedApplicationEvents"`

	// IpsEvents - no documentation
	IpsEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"ipsEvents"`

	// ApplicationModePolicyNames - The names of the possible policy options for the application mode
	// setting.
	ApplicationModePolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"applicationModePolicyNames"`

	// FirewallModePolicyNames - The names of the possible policy options for the firewall mode setting.
	FirewallModePolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"firewallModePolicyNames"`

	// FirewallRuleSetPolicyNames - The names of the possible policy options for the firewall rule set
	// setting.
	FirewallRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"firewallRuleSetPolicyNames"`

	// IpsModePolicyNames - The names of the possible policy options for the host IPS mode setting.
	IpsModePolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"ipsModePolicyNames"`

	// ApplicationRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// application rule set setting.
	ApplicationRuleSetPolicyNameCount uint64 `json:"applicationRuleSetPolicyNameCount"`

	// FirewallModePolicyNameCount - A count of the names of the possible policy options for the firewall
	// mode setting.
	FirewallModePolicyNameCount uint64 `json:"firewallModePolicyNameCount"`

	// IpsProtectionPolicyNameCount - A count of the names of the possible policy options for the host IPS
	// protection setting.
	IpsProtectionPolicyNameCount uint64 `json:"ipsProtectionPolicyNameCount"`
}

func (softlayer_software_component_hostips_mcafee_epo_version45_hips *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips"
}
