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

	// PrimarySubnetId - The internal identifier of the primary subnet addressed on a
	PrimarySubnetId int `json:"primarySubnetId"`

	// Name - no documentation
	Name string `json:"name"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Id - A VLAN's internal identifier. This should not be confused with the ''vlanNumber'' property,
	// which is used in network configuration.
	Id int `json:"id"`

	// Note - no documentation
	Note string `json:"note"`

	// VlanNumber - A VLAN's number as recorded within the SoftLayer network. This is configured directly
	// on SoftLayer's networking equipment and should not be confused with a VLAN's ''id'' property.
	VlanNumber int `json:"vlanNumber"`

	// AccountId - The internal identifier of the SoftLayer customer account that a is associated with.
	AccountId int `json:"accountId"`
}

// SoftLayer_Network_Vlan_Extended is SoftLayer_Network_Vlan with all maskable types.
type SoftLayer_Network_Vlan_Extended struct {
	SoftLayer_Network_Vlan

	// HighAvailabilityFirewallFlag - <nil>
	HighAvailabilityFirewallFlag bool `json:"highAvailabilityFirewallFlag"`

	// LocalDiskStorageCapabilityFlag - A flag indicating that a vlan can be assigned to a host that has
	// local disk functionality.
	LocalDiskStorageCapabilityFlag bool `json:"localDiskStorageCapabilityFlag"`

	// FirewallRuleCount - A count of the currently running rule set of a firewalled
	FirewallRuleCount uint64 `json:"firewallRuleCount"`

	// PrivateNetworkGatewayCount - no documentation
	PrivateNetworkGatewayCount uint64 `json:"privateNetworkGatewayCount"`

	// NetworkVlanFirewall - The Hardware Firewall (Dedicated) for a network vlan.
	NetworkVlanFirewall *SoftLayer_Network_Vlan_Firewall `json:"networkVlanFirewall"`

	// PrivateNetworkGateways - no documentation
	PrivateNetworkGateways []*SoftLayer_Network_Gateway `json:"privateNetworkGateways"`

	// FirewallInterfaceCount - A count of a firewalled vlan's inbound/outbound interfaces.
	FirewallInterfaceCount uint64 `json:"firewallInterfaceCount"`

	// FirewallNetworkComponentCount - no documentation
	FirewallNetworkComponentCount uint64 `json:"firewallNetworkComponentCount"`

	// HardwareCount - A count of all of the hardware that exists on a Hardware is associated with a by its
	// networking components.
	HardwareCount uint64 `json:"hardwareCount"`

	// AttachedNetworkGatewayVlan - The inside record if this is inside a network gateway.
	AttachedNetworkGatewayVlan *SoftLayer_Network_Gateway_Vlan `json:"attachedNetworkGatewayVlan"`

	// NetworkComponentTrunks - The network components that are connected to this through a trunk.
	NetworkComponentTrunks []*SoftLayer_Network_Component_Network_Vlan_Trunk `json:"networkComponentTrunks"`

	// SecondarySubnets - The subnets that exist as secondary interfaces on a
	SecondarySubnets []*SoftLayer_Network_Subnet `json:"secondarySubnets"`

	// PrimarySubnetCount - no documentation
	PrimarySubnetCount uint64 `json:"primarySubnetCount"`

	// NetworkComponents - no documentation
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents"`

	// NetworkSpace - Identifier to denote whether a is used for public or private connectivity.
	NetworkSpace string `json:"networkSpace"`

	// SanStorageCapabilityFlag - A flag indicating that a vlan can be assigned to a host that has SAN disk
	// functionality.
	SanStorageCapabilityFlag bool `json:"sanStorageCapabilityFlag"`

	// AdditionalPrimarySubnetCount - A count of a VLAN's additional primary subnets. These are used to
	// extend the number of servers attached to the by adding more ip addresses to the primary IP address
	// pool.
	AdditionalPrimarySubnetCount uint64 `json:"additionalPrimarySubnetCount"`

	// SecondarySubnetCount - A count of the subnets that exist as secondary interfaces on a
	SecondarySubnetCount uint64 `json:"secondarySubnetCount"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// DedicatedFirewallFlag - A flag indicating that a network vlan is on a Hardware Firewall (Dedicated).
	DedicatedFirewallFlag int `json:"dedicatedFirewallFlag"`

	// ExtensionRouter - no documentation
	ExtensionRouter *SoftLayer_Hardware_Router `json:"extensionRouter"`

	// FirewallNetworkComponents - no documentation
	FirewallNetworkComponents []*SoftLayer_Network_Component_Firewall `json:"firewallNetworkComponents"`

	// PublicNetworkGateways - no documentation
	PublicNetworkGateways []*SoftLayer_Network_Gateway `json:"publicNetworkGateways"`

	// ProtectedIpAddresses - <nil>
	ProtectedIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"protectedIpAddresses"`

	// NetworkComponentTrunkCount - A count of the network components that are connected to this through a
	// trunk.
	NetworkComponentTrunkCount uint64 `json:"networkComponentTrunkCount"`

	// PublicNetworkGatewayCount - no documentation
	PublicNetworkGatewayCount uint64 `json:"publicNetworkGatewayCount"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount"`

	// AdditionalPrimarySubnets - A VLAN's additional primary subnets. These are used to extend the number
	// of servers attached to the by adding more ip addresses to the primary IP address pool.
	AdditionalPrimarySubnets []*SoftLayer_Network_Subnet `json:"additionalPrimarySubnets"`

	// FirewallInterfaces - no documentation
	FirewallInterfaces []*SoftLayer_Network_Firewall_Module_Context_Interface `json:"firewallInterfaces"`

	// PrimarySubnet - A VLAN's primary subnet. Each has at least one subnet, usually the subnet that is
	// assigned to a server or new IP address block when it's purchased.
	PrimarySubnet *SoftLayer_Network_Subnet `json:"primarySubnet"`

	// Type - no documentation
	Type *SoftLayer_Network_Vlan_Type `json:"type"`

	// VirtualGuests - All of the Virtual Servers that are connected to a
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`

	// ResourceGroupMember - no documentation
	ResourceGroupMember []*SoftLayer_Resource_Group_Member `json:"resourceGroupMember"`

	// FirewallGuestNetworkComponentCount - no documentation
	FirewallGuestNetworkComponentCount uint64 `json:"firewallGuestNetworkComponentCount"`

	// NetworkComponentCount - A count of the networking components that are connected to a
	NetworkComponentCount uint64 `json:"networkComponentCount"`

	// SubnetCount - A count of all of the subnets that exist as interfaces.
	SubnetCount uint64 `json:"subnetCount"`

	// VirtualGuestCount - A count of all of the Virtual Servers that are connected to a
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// AttachedNetworkGatewayFlag - no documentation
	AttachedNetworkGatewayFlag bool `json:"attachedNetworkGatewayFlag"`

	// ResourceGroupCount - A count of the resource groups in which this is a member.
	ResourceGroupCount uint64 `json:"resourceGroupCount"`

	// AttachedNetworkGateway - no documentation
	AttachedNetworkGateway *SoftLayer_Network_Gateway `json:"attachedNetworkGateway"`

	// GuestNetworkComponents - no documentation
	GuestNetworkComponents []*SoftLayer_Virtual_Guest_Network_Component `json:"guestNetworkComponents"`

	// Hardware - All of the hardware that exists on a Hardware is associated with a by its networking
	// components.
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// TotalPrimaryIpAddressCount - no documentation
	TotalPrimaryIpAddressCount uint `json:"totalPrimaryIpAddressCount"`

	// PrimarySubnetVersion6 - A VLAN's primary IPv6 subnet. Some VLAN's may not have a primary IPv6
	// subnet.
	PrimarySubnetVersion6 *SoftLayer_Network_Subnet `json:"primarySubnetVersion6"`

	// ResourceGroups - no documentation
	ResourceGroups []*SoftLayer_Resource_Group `json:"resourceGroups"`

	// SecondaryRouter - The secondary router that a is associated with. Every SoftLayer is connected to
	// more than one router for greater network redundancy.
	SecondaryRouter *SoftLayer_Hardware `json:"secondaryRouter"`

	// ProtectedIpAddressCount - no documentation
	ProtectedIpAddressCount uint64 `json:"protectedIpAddressCount"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// FirewallRules - no documentation
	FirewallRules []*SoftLayer_Network_Vlan_Firewall_Rule `json:"firewallRules"`

	// ScaleVlans - no documentation
	ScaleVlans []*SoftLayer_Scale_Network_Vlan `json:"scaleVlans"`

	// ResourceGroupMemberCount - A count of the resource group member for a network vlan.
	ResourceGroupMemberCount uint64 `json:"resourceGroupMemberCount"`

	// ScaleVlanCount - A count of collection of scale VLANs this applies to.
	ScaleVlanCount uint64 `json:"scaleVlanCount"`

	// PrimarySubnets - <nil>
	PrimarySubnets []*SoftLayer_Network_Subnet `json:"primarySubnets"`

	// Subnets - no documentation
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// FirewallGuestNetworkComponents - no documentation
	FirewallGuestNetworkComponents []*SoftLayer_Network_Component_Firewall `json:"firewallGuestNetworkComponents"`

	// PrimaryRouter - The primary router that a is associated with. Every SoftLayer is connected to more
	// than one router for greater network redundancy.
	PrimaryRouter *SoftLayer_Hardware_Router `json:"primaryRouter"`

	// TagReferences - no documentation
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences"`

	// GuestNetworkComponentCount - A count of the networking components that are connected to a
	GuestNetworkComponentCount uint64 `json:"guestNetworkComponentCount"`
}

func (softlayer_network_vlan *SoftLayer_Network_Vlan) String() string {
	return "SoftLayer_Network_Vlan"
}
