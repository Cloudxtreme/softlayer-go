package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 - The
// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 data type represents a
// single McAfee Secure Host IPS software component for version 6 of the Host IPS client and uses the
// ePolicy Orchestrator version 3.6 backend.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 struct {

	// ManufacturerLicenseInstance - A license key for this specific installation of software, if it is
	// needed.
	ManufacturerLicenseInstance string `json:"manufacturerLicenseInstance,omitempty"`

	// Id - An ID number identifying this Software Component (Software Installation)
	Id int `json:"id,omitempty"`

	// EpoVersion - The version of ePolicy Orchestrator that the host IPS client communicates with.
	EpoVersion string `json:"epoVersion,omitempty"`

	// HardwareId - Hardware Identification Number for the server this Software Component is installed
	// upon.
	HardwareId int `json:"hardwareId,omitempty"`

	// ManufacturerActivationCode - The manufacturer code that is needed to activate a license.
	ManufacturerActivationCode string `json:"manufacturerActivationCode,omitempty"`
}

func (softlayer_software_component_hostips_mcafee_epo_version36_hips_version6 *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6"
}

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6_Extended is SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 with all maskable types.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6_Extended struct {
	SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6

	// IpsEvents - no documentation
	IpsEvents []*McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent `json:"ipsEvents,omitempty"`

	// IpsEventCount - A count of the host IPS events for this software component.
	IpsEventCount uint64 `json:"ipsEventCount,omitempty"`

	// PasswordHistoryCount - no documentation
	PasswordHistoryCount uint64 `json:"passwordHistoryCount,omitempty"`

	// ApplicationModePolicyNameCount - A count of the names of the possible policy options for the
	// application mode setting.
	ApplicationModePolicyNameCount uint64 `json:"applicationModePolicyNameCount,omitempty"`

	// AgentDetails - no documentation
	AgentDetails *McAfee_Epolicy_Orchestrator_Version36_Agent_Details `json:"agentDetails,omitempty"`

	// BlockedApplicationEvents - The blocked application events for this software component.
	BlockedApplicationEvents []*McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent `json:"blockedApplicationEvents,omitempty"`

	// ApplicationRuleSetPolicyNames - The names of the possible policy options for the application rule
	// set setting.
	ApplicationRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationRuleSetPolicyNames,omitempty"`

	// TransactionStatus - no documentation
	TransactionStatus string `json:"transactionStatus,omitempty"`

	// SoftwareLicense - no documentation
	SoftwareLicense *SoftLayer_Software_License `json:"softwareLicense,omitempty"`

	// EnforcementPolicyNameCount - A count of the names of the possible options for the enforcement policy
	// setting.
	EnforcementPolicyNameCount uint64 `json:"enforcementPolicyNameCount,omitempty"`

	// ApplicationRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// application rule set setting.
	ApplicationRuleSetPolicyNameCount uint64 `json:"applicationRuleSetPolicyNameCount,omitempty"`

	// ApplicationModePolicyNames - The names of the possible policy options for the application mode
	// setting.
	ApplicationModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"applicationModePolicyNames,omitempty"`

	// IpsModePolicyNames - The names of the possible policy options for the host IPS mode setting.
	IpsModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsModePolicyNames,omitempty"`

	// BlockedApplicationEventCount - A count of the blocked application events for this software
	// component.
	BlockedApplicationEventCount uint64 `json:"blockedApplicationEventCount,omitempty"`

	// Hardware - The hardware this Software Component is installed upon.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// FirewallModePolicyNames - The names of the possible policy options for the firewall mode setting.
	FirewallModePolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallModePolicyNames,omitempty"`

	// FirewallRuleSetPolicyNames - The names of the possible policy options for the firewall rule set
	// setting.
	FirewallRuleSetPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"firewallRuleSetPolicyNames,omitempty"`

	// EnforcementPolicyNames - The names of the possible options for the enforcement policy setting.
	EnforcementPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"enforcementPolicyNames,omitempty"`

	// IpsProtectionPolicyNameCount - A count of the names of the possible policy options for the host IPS
	// protection setting.
	IpsProtectionPolicyNameCount uint64 `json:"ipsProtectionPolicyNameCount,omitempty"`

	// AverageInstallationDuration - The average amount of time that a software component takes to install.
	AverageInstallationDuration uint64 `json:"averageInstallationDuration,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// SoftwareDescription - The Software Description of this Software Component.
	SoftwareDescription *SoftLayer_Software_Description `json:"softwareDescription,omitempty"`

	// FirewallRuleSetPolicyNameCount - A count of the names of the possible policy options for the
	// firewall rule set setting.
	FirewallRuleSetPolicyNameCount uint64 `json:"firewallRuleSetPolicyNameCount,omitempty"`

	// FirewallModePolicyNameCount - A count of the names of the possible policy options for the firewall
	// mode setting.
	FirewallModePolicyNameCount uint64 `json:"firewallModePolicyNameCount,omitempty"`

	// Passwords - Username/Password pairs used for access to this Software Installation.
	Passwords []*SoftLayer_Software_Component_Password `json:"passwords,omitempty"`

	// IpsModePolicyNameCount - A count of the names of the possible policy options for the host IPS mode
	// setting.
	IpsModePolicyNameCount uint64 `json:"ipsModePolicyNameCount,omitempty"`

	// IpsProtectionPolicyNames - The names of the possible policy options for the host IPS protection
	// setting.
	IpsProtectionPolicyNames []*McAfee_Epolicy_Orchestrator_Version36_Policy_Object `json:"ipsProtectionPolicyNames,omitempty"`

	// PasswordHistory - no documentation
	PasswordHistory []*SoftLayer_Software_Component_Password_History `json:"passwordHistory,omitempty"`

	// PasswordCount - A count of username/Password pairs used for access to this Software Installation.
	PasswordCount uint64 `json:"passwordCount,omitempty"`

	// VirtualGuest - The virtual guest this software component is installed upon.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`
}

func (softlayer_software_component_hostips_mcafee_epo_version36_hips_version6 *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6_Extended) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6"
}
