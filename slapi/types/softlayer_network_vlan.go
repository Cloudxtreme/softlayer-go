package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Vlan - The SoftLayer_Network_Vlan data type models a single within SoftLayer's
// public and private networks. a Virtual LAN is a structure that associates network interfaces on
// routers, switches, and servers in different locations to act as if they were on the same local
// network broadcast domain. VLANs are a central part of the SoftLayer network. They can determine how
// new IP subnets are routed and how individual servers communicate to each other.
type SoftLayer_Network_Vlan struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The internal identifier of the SoftLayer customer account that a is associated with.
	AccountId int `json:"accountId"`

	// AdditionalPrimarySubnetCount - A count of a VLAN's additional primary subnets. These are used to
	// extend the number of servers attached to the by adding more ip addresses to the primary IP address
	// pool.
	AdditionalPrimarySubnetCount uint64 `json:"additionalPrimarySubnetCount"`

	// AdditionalPrimarySubnets - A VLAN's additional primary subnets. These are used to extend the number
	// of servers attached to the by adding more ip addresses to the primary IP address pool.
	AdditionalPrimarySubnets []*SoftLayer_Network_Subnet `json:"additionalPrimarySubnets"`

	// AttachedNetworkGateway - no documentation
	AttachedNetworkGateway *SoftLayer_Network_Gateway `json:"attachedNetworkGateway"`

	// AttachedNetworkGatewayFlag - no documentation
	AttachedNetworkGatewayFlag bool `json:"attachedNetworkGatewayFlag"`

	// AttachedNetworkGatewayVlan - The inside record if this is inside a network gateway.
	AttachedNetworkGatewayVlan *SoftLayer_Network_Gateway_Vlan `json:"attachedNetworkGatewayVlan"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// DedicatedFirewallFlag - A flag indicating that a network vlan is on a Hardware Firewall (Dedicated).
	DedicatedFirewallFlag int `json:"dedicatedFirewallFlag"`

	// ExtensionRouter - no documentation
	ExtensionRouter *SoftLayer_Hardware_Router `json:"extensionRouter"`

	// FirewallGuestNetworkComponentCount - no documentation
	FirewallGuestNetworkComponentCount uint64 `json:"firewallGuestNetworkComponentCount"`

	// FirewallGuestNetworkComponents - no documentation
	FirewallGuestNetworkComponents []*SoftLayer_Network_Component_Firewall `json:"firewallGuestNetworkComponents"`

	// FirewallInterfaceCount - A count of a firewalled vlan's inbound/outbound interfaces.
	FirewallInterfaceCount uint64 `json:"firewallInterfaceCount"`

	// FirewallInterfaces - no documentation
	FirewallInterfaces []*SoftLayer_Network_Firewall_Module_Context_Interface `json:"firewallInterfaces"`

	// FirewallNetworkComponentCount - no documentation
	FirewallNetworkComponentCount uint64 `json:"firewallNetworkComponentCount"`

	// FirewallNetworkComponents - no documentation
	FirewallNetworkComponents []*SoftLayer_Network_Component_Firewall `json:"firewallNetworkComponents"`

	// FirewallRuleCount - A count of the currently running rule set of a firewalled
	FirewallRuleCount uint64 `json:"firewallRuleCount"`

	// FirewallRules - no documentation
	FirewallRules []*SoftLayer_Network_Vlan_Firewall_Rule `json:"firewallRules"`

	// GuestNetworkComponentCount - A count of the networking components that are connected to a
	GuestNetworkComponentCount uint64 `json:"guestNetworkComponentCount"`

	// GuestNetworkComponents - no documentation
	GuestNetworkComponents []*SoftLayer_Virtual_Guest_Network_Component `json:"guestNetworkComponents"`

	// Hardware - All of the hardware that exists on a Hardware is associated with a by its networking
	// components.
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// HardwareCount - A count of all of the hardware that exists on a Hardware is associated with a by its
	// networking components.
	HardwareCount uint64 `json:"hardwareCount"`

	// HighAvailabilityFirewallFlag - <nil>
	HighAvailabilityFirewallFlag bool `json:"highAvailabilityFirewallFlag"`

	// Id - A VLAN's internal identifier. This should not be confused with the ''vlanNumber'' property,
	// which is used in network configuration.
	Id int `json:"id"`

	// LocalDiskStorageCapabilityFlag - A flag indicating that a vlan can be assigned to a host that has
	// local disk functionality.
	LocalDiskStorageCapabilityFlag bool `json:"localDiskStorageCapabilityFlag"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// NetworkComponentCount - A count of the networking components that are connected to a
	NetworkComponentCount uint64 `json:"networkComponentCount"`

	// NetworkComponentTrunkCount - A count of the network components that are connected to this through a
	// trunk.
	NetworkComponentTrunkCount uint64 `json:"networkComponentTrunkCount"`

	// NetworkComponentTrunks - The network components that are connected to this through a trunk.
	NetworkComponentTrunks []*SoftLayer_Network_Component_Network_Vlan_Trunk `json:"networkComponentTrunks"`

	// NetworkComponents - no documentation
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents"`

	// NetworkSpace - Identifier to denote whether a is used for public or private connectivity.
	NetworkSpace string `json:"networkSpace"`

	// NetworkVlanFirewall - The Hardware Firewall (Dedicated) for a network vlan.
	NetworkVlanFirewall *SoftLayer_Network_Vlan_Firewall `json:"networkVlanFirewall"`

	// Note - no documentation
	Note string `json:"note"`

	// PrimaryRouter - The primary router that a is associated with. Every SoftLayer is connected to more
	// than one router for greater network redundancy.
	PrimaryRouter *SoftLayer_Hardware_Router `json:"primaryRouter"`

	// PrimarySubnet - A VLAN's primary subnet. Each has at least one subnet, usually the subnet that is
	// assigned to a server or new IP address block when it's purchased.
	PrimarySubnet *SoftLayer_Network_Subnet `json:"primarySubnet"`

	// PrimarySubnetCount - no documentation
	PrimarySubnetCount uint64 `json:"primarySubnetCount"`

	// PrimarySubnetId - The internal identifier of the primary subnet addressed on a
	PrimarySubnetId int `json:"primarySubnetId"`

	// PrimarySubnetVersion6 - A VLAN's primary IPv6 subnet. Some VLAN's may not have a primary IPv6
	// subnet.
	PrimarySubnetVersion6 *SoftLayer_Network_Subnet `json:"primarySubnetVersion6"`

	// PrimarySubnets - <nil>
	PrimarySubnets []*SoftLayer_Network_Subnet `json:"primarySubnets"`

	// PrivateNetworkGatewayCount - no documentation
	PrivateNetworkGatewayCount uint64 `json:"privateNetworkGatewayCount"`

	// PrivateNetworkGateways - no documentation
	PrivateNetworkGateways []*SoftLayer_Network_Gateway `json:"privateNetworkGateways"`

	// ProtectedIpAddressCount - no documentation
	ProtectedIpAddressCount uint64 `json:"protectedIpAddressCount"`

	// ProtectedIpAddresses - <nil>
	ProtectedIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"protectedIpAddresses"`

	// PublicNetworkGatewayCount - no documentation
	PublicNetworkGatewayCount uint64 `json:"publicNetworkGatewayCount"`

	// PublicNetworkGateways - no documentation
	PublicNetworkGateways []*SoftLayer_Network_Gateway `json:"publicNetworkGateways"`

	// ResourceGroupCount - A count of the resource groups in which this is a member.
	ResourceGroupCount uint64 `json:"resourceGroupCount"`

	// ResourceGroupMember - no documentation
	ResourceGroupMember []*SoftLayer_Resource_Group_Member `json:"resourceGroupMember"`

	// ResourceGroupMemberCount - A count of the resource group member for a network vlan.
	ResourceGroupMemberCount uint64 `json:"resourceGroupMemberCount"`

	// ResourceGroups - no documentation
	ResourceGroups []*SoftLayer_Resource_Group `json:"resourceGroups"`

	// SanStorageCapabilityFlag - A flag indicating that a vlan can be assigned to a host that has SAN disk
	// functionality.
	SanStorageCapabilityFlag bool `json:"sanStorageCapabilityFlag"`

	// ScaleVlanCount - A count of collection of scale VLANs this applies to.
	ScaleVlanCount uint64 `json:"scaleVlanCount"`

	// ScaleVlans - no documentation
	ScaleVlans []*SoftLayer_Scale_Network_Vlan `json:"scaleVlans"`

	// SecondaryRouter - The secondary router that a is associated with. Every SoftLayer is connected to
	// more than one router for greater network redundancy.
	SecondaryRouter *SoftLayer_Hardware `json:"secondaryRouter"`

	// SecondarySubnetCount - A count of the subnets that exist as secondary interfaces on a
	SecondarySubnetCount uint64 `json:"secondarySubnetCount"`

	// SecondarySubnets - The subnets that exist as secondary interfaces on a
	SecondarySubnets []*SoftLayer_Network_Subnet `json:"secondarySubnets"`

	// SubnetCount - A count of all of the subnets that exist as interfaces.
	SubnetCount uint64 `json:"subnetCount"`

	// Subnets - no documentation
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount"`

	// TagReferences - no documentation
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences"`

	// TotalPrimaryIpAddressCount - no documentation
	TotalPrimaryIpAddressCount uint `json:"totalPrimaryIpAddressCount"`

	// Type - no documentation
	Type *SoftLayer_Network_Vlan_Type `json:"type"`

	// VirtualGuestCount - A count of all of the Virtual Servers that are connected to a
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// VirtualGuests - All of the Virtual Servers that are connected to a
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`

	// VlanNumber - A VLAN's number as recorded within the SoftLayer network. This is configured directly
	// on SoftLayer's networking equipment and should not be confused with a VLAN's ''id'' property.
	VlanNumber int `json:"vlanNumber"`
}

func (softlayer_network_vlan *SoftLayer_Network_Vlan) String() string {
	return "SoftLayer_Network_Vlan"
}
