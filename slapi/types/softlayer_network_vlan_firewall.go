package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Vlan_Firewall - The SoftLayer_Network_Vlan_Firewall data type contains general
// information relating to a single SoftLayer firewall. This is the object which ties the running rules
// to a specific downstream server. Use the [[SoftLayer Network Firewall Template]] service to pull
// SoftLayer recommended rule set templates. Use the [[SoftLayer Network Firewall Update Request]]
// service to submit a firewall update request.
type SoftLayer_Network_Vlan_Firewall struct {

	// CustomerManagedFlag - Whether or not this firewall can be directly logged in to.
	CustomerManagedFlag bool `json:"customerManagedFlag,omitempty"`

	// PrimaryIpAddress - A firewall's primary IP address. This field will be the IP shown when doing
	// network traces and reverse DNS and is a read-only property.
	PrimaryIpAddress string `json:"primaryIpAddress,omitempty"`

	// AdministrativeBypassFlag - A flag to indicate if the firewall is in administrative bypass mode. In
	// other words, no rules are being applied to the traffic coming through.
	AdministrativeBypassFlag string `json:"administrativeBypassFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// FullyQualifiedDomainName - A name reflecting the hostname and domain of the firewall. This is
	// created from the combined values of the firewall's logical name and vlan number automatically, and
	// thus can not be edited directly.
	FullyQualifiedDomainName string `json:"fullyQualifiedDomainName,omitempty"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount,omitempty"`

	// NetworkVlan - The object that a firewall is associated with and protecting.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan,omitempty"`

	// NetworkVlans - The objects that a firewall is associated with and protecting.
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans,omitempty"`

	// NetworkVlanCount - A count of the objects that a firewall is associated with and protecting.
	NetworkVlanCount uint64 `json:"networkVlanCount,omitempty"`

	// RuleCount - A count of the currently running rule set of this network component firewall.
	RuleCount uint64 `json:"ruleCount,omitempty"`

	// Rules - The currently running rule set of this network component firewall.
	Rules []*SoftLayer_Network_Vlan_Firewall_Rule `json:"rules,omitempty"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences,omitempty"`

	// BillingItem - The billing item for a Hardware Firewall (Dedicated).
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// FirewallType - no documentation
	FirewallType string `json:"firewallType,omitempty"`

	// ManagementCredentials - The credentials to log in to a firewall device. This is only present for
	// dedicated appliances.
	ManagementCredentials *SoftLayer_Software_Component_Password `json:"managementCredentials,omitempty"`

	// NetworkFirewallUpdateRequests - no documentation
	NetworkFirewallUpdateRequests []*SoftLayer_Network_Firewall_Update_Request `json:"networkFirewallUpdateRequests,omitempty"`

	// NetworkFirewallUpdateRequestCount - A count of the update requests made for this firewall.
	NetworkFirewallUpdateRequestCount uint64 `json:"networkFirewallUpdateRequestCount,omitempty"`
}

func (softlayer_network_vlan_firewall *SoftLayer_Network_Vlan_Firewall) String() string {
	return "SoftLayer_Network_Vlan_Firewall"
}
