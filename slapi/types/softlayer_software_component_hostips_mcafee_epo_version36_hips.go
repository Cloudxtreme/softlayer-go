package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips - The
// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips data type represents a single McAfee
// Secure Host IPS software component that uses the ePolicy Orchestrator version 3.6 backend.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips struct {

	// Id - An ID number identifying this Software Component (Software Installation)
	Id int `json:"id,omitempty"`

	// ManufacturerLicenseInstance - A license key for this specific installation of software, if it is
	// needed.
	ManufacturerLicenseInstance string `json:"manufacturerLicenseInstance,omitempty"`

	// HardwareId - Hardware Identification Number for the server this Software Component is installed
	// upon.
	HardwareId int `json:"hardwareId,omitempty"`

	// ManufacturerActivationCode - The manufacturer code that is needed to activate a license.
	ManufacturerActivationCode string `json:"manufacturerActivationCode,omitempty"`

	// EpoVersion - The version of ePolicy Orchestrator that the host IPS client communicates with.
	EpoVersion string `json:"epoVersion,omitempty"`

	// SoftwareDescription - The Software Description of this Software Component.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// ApplicationRuleSetPolicyNames - The names of the possible policy options for the application rule
	// set setting.
	ApplicationRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationRuleSetPolicyNames,omitempty"`

	// EnforcementPolicyNames - The names of the possible options for the enforcement policy setting.
	EnforcementPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"enforcementPolicyNames,omitempty"`

	// FirewallRuleSetPolicyNames - The names of the possible policy options for the firewall rule set
	// setting.
	FirewallRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallRuleSetPolicyNames,omitempty"`

	// IpsProtectionPolicyNames - The names of the possible policy options for the host IPS protection
	// setting.
	IpsProtectionPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsProtectionPolicyNames,omitempty"`

	// ApplicationRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// application rule set setting.
	ApplicationRuleSetPolicyNameCount uint64 `json:"applicationRuleSetPolicyNameCount,omitempty"`

	// FirewallModePolicyNameCount - A count of the names of the possible policy options for the firewall
	// mode setting.
	FirewallModePolicyNameCount uint64 `json:"firewallModePolicyNameCount,omitempty"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version36_Agent_Details `json:"agentDetails,omitempty"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus,omitempty"`

	// PasswordHistoryCount - no documentation
	PasswordHistoryCount uint64 `json:"passwordHistoryCount,omitempty"`

	// PasswordHistory - no documentation
	PasswordHistory []*SoftLayer_Software_Component_Password_History `json:"passwordHistory,omitempty"`

	// SoftwareLicense - no documentation
	SoftwareLicense *SoftLayer_Software_License `json:"softwareLicense,omitempty"`

	// VirtualGuest - The virtual guest this software component is installed upon.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// IpsModePolicyNames - The names of the possible policy options for the host IPS mode setting.
	IpsModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsModePolicyNames,omitempty"`

	// IpsModePolicyNameCount - A count of the names of the possible policy options for the host IPS mode
	// setting.
	IpsModePolicyNameCount uint64 `json:"ipsModePolicyNameCount,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// ApplicationModePolicyNameCount - A count of the names of the possible policy options for the
	// application mode setting.
	ApplicationModePolicyNameCount uint64 `json:"applicationModePolicyNameCount,omitempty"`

	// FirewallModePolicyNames - The names of the possible policy options for the firewall mode setting.
	FirewallModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallModePolicyNames,omitempty"`

	// EnforcementPolicyNameCount - A count of the names of the possible options for the enforcement policy
	// setting.
	EnforcementPolicyNameCount uint64 `json:"enforcementPolicyNameCount,omitempty"`

	// Passwords - Username/Password pairs used for access to this Software Installation.
	Passwords []*SoftLayer_Software_Component_Password `json:"passwords,omitempty"`

	// FirewallRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// firewall rule set setting.
	FirewallRuleSetPolicyNameCount uint64 `json:"firewallRuleSetPolicyNameCount,omitempty"`

	// PasswordCount - A count of username/Password pairs used for access to this Software Installation.
	PasswordCount uint64 `json:"passwordCount,omitempty"`

	// Hardware - The hardware this Software Component is installed upon.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// ApplicationModePolicyNames - The names of the possible policy options for the application mode
	// setting.
	ApplicationModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationModePolicyNames,omitempty"`

	// IpsProtectionPolicyNameCount - A count of the names of the possible policy options for the host IPS
	// protection setting.
	IpsProtectionPolicyNameCount uint64 `json:"ipsProtectionPolicyNameCount,omitempty"`

	// AverageInstallationDuration - The average amount of time that a software component takes to install.
	AverageInstallationDuration uint64 `json:"averageInstallationDuration,omitempty"`
}

func (softlayer_software_component_hostips_mcafee_epo_version36_hips *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips"
}
