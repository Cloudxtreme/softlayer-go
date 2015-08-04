package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips - The
// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips data type represents a single McAfee
// Secure Host IPS software component that uses the ePolicy Orchestrator version 4.5 backend.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips struct {

	// EpoVersion - The version of ePolicy Orchestrator that the host IPS client communicates with.
	EpoVersion string `json:"epoVersion,omitempty"`

	// HardwareId - Hardware Identification Number for the server this Software Component is installed
	// upon.
	HardwareId int `json:"hardwareId,omitempty"`

	// ManufacturerActivationCode - The manufacturer code that is needed to activate a license.
	ManufacturerActivationCode string `json:"manufacturerActivationCode,omitempty"`

	// ManufacturerLicenseInstance - A license key for this specific installation of software, if it is
	// needed.
	ManufacturerLicenseInstance string `json:"manufacturerLicenseInstance,omitempty"`

	// Id - An ID number identifying this Software Component (Software Installation)
	Id int `json:"id,omitempty"`
}

func (softlayer_software_component_hostips_mcafee_epo_version45_hips *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips"
}

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Extended is SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips with all maskable types.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Extended struct {
	SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips

	// Hardware - The hardware this Software Component is installed upon.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// SoftwareLicense - no documentation
	SoftwareLicense *SoftLayer_Software_License `json:"softwareLicense,omitempty"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version45_Agent_Details `json:"agentDetails,omitempty"`

	// FirewallRuleSetPolicyNames - The names of the possible policy options for the firewall rule set
	// setting.
	FirewallRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"firewallRuleSetPolicyNames,omitempty"`

	// ApplicationRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// application rule set setting.
	ApplicationRuleSetPolicyNameCount uint64 `json:"applicationRuleSetPolicyNameCount,omitempty"`

	// VirtualGuest - The virtual guest this software component is installed upon.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// IpsEvents - no documentation
	IpsEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"ipsEvents,omitempty"`

	// IpsModePolicyNames - The names of the possible policy options for the host IPS mode setting.
	IpsModePolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"ipsModePolicyNames,omitempty"`

	// BlockedApplicationEventCount - A count of the blocked application events for this software
	// component.
	BlockedApplicationEventCount uint64 `json:"blockedApplicationEventCount,omitempty"`

	// IpsProtectionPolicyNameCount - A count of the names of the possible policy options for the host IPS
	// protection setting.
	IpsProtectionPolicyNameCount uint64 `json:"ipsProtectionPolicyNameCount,omitempty"`

	// SoftwareDescription - The Software Description of this Software Component.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// PasswordHistoryCount - no documentation
	PasswordHistoryCount uint64 `json:"passwordHistoryCount,omitempty"`

	// Passwords - Username/Password pairs used for access to this Software Installation.
	Passwords []*SoftLayer_Software_Component_Password `json:"passwords,omitempty"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus,omitempty"`

	// EnforcementPolicyNameCount - A count of the names of the possible options for the enforcement policy
	// setting.
	EnforcementPolicyNameCount uint64 `json:"enforcementPolicyNameCount,omitempty"`

	// IpsEventCount - A count of the host IPS events for this software component.
	IpsEventCount uint64 `json:"ipsEventCount,omitempty"`

	// IpsModePolicyNameCount - A count of the names of the possible policy options for the host IPS mode
	// setting.
	IpsModePolicyNameCount uint64 `json:"ipsModePolicyNameCount,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// ApplicationRuleSetPolicyNames - The names of the possible policy options for the application rule
	// set setting.
	ApplicationRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"applicationRuleSetPolicyNames,omitempty"`

	// EnforcementPolicyNames - The names of the possible options for the enforcement policy setting.
	EnforcementPolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"enforcementPolicyNames,omitempty"`

	// FirewallRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// firewall rule set setting.
	FirewallRuleSetPolicyNameCount uint64 `json:"firewallRuleSetPolicyNameCount,omitempty"`

	// PasswordCount - A count of username/Password pairs used for access to this Software Installation.
	PasswordCount uint64 `json:"passwordCount,omitempty"`

	// ApplicationModePolicyNames - The names of the possible policy options for the application mode
	// setting.
	ApplicationModePolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"applicationModePolicyNames,omitempty"`

	// IpsProtectionPolicyNames - The names of the possible policy options for the host IPS protection
	// setting.
	IpsProtectionPolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"ipsProtectionPolicyNames,omitempty"`

	// FirewallModePolicyNameCount - A count of the names of the possible policy options for the firewall
	// mode setting.
	FirewallModePolicyNameCount uint64 `json:"firewallModePolicyNameCount,omitempty"`

	// BlockedApplicationEvents - The blocked application events for this software component.
	BlockedApplicationEvents []*McAfee_Epolicy_Orchestrator_Version45_Event `json:"blockedApplicationEvents,omitempty"`

	// FirewallModePolicyNames - The names of the possible policy options for the firewall mode setting.
	FirewallModePolicyNames []*McAfee_Epolicy_Orchestrator_Version45_Policy_Object `json:"firewallModePolicyNames,omitempty"`

	// ApplicationModePolicyNameCount - A count of the names of the possible policy options for the
	// application mode setting.
	ApplicationModePolicyNameCount uint64 `json:"applicationModePolicyNameCount,omitempty"`

	// PasswordHistory - no documentation
	PasswordHistory []*SoftLayer_Software_Component_Password_History `json:"passwordHistory,omitempty"`

	// AverageInstallationDuration - The average amount of time that a software component takes to install.
	AverageInstallationDuration uint64 `json:"averageInstallationDuration,omitempty"`
}

func (softlayer_software_component_hostips_mcafee_epo_version45_hips *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips_Extended) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version45_Hips"
}
