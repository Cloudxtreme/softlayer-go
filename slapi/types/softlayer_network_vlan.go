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

	// Note - no documentation
	Note string `json:"note,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Id - A VLAN's internal identifier. This should not be confused with the ''vlanNumber'' property,
	// which is used in network configuration.
	Id int `json:"id,omitempty"`

	// PrimarySubnetId - The internal identifier of the primary subnet addressed on a
	PrimarySubnetId int `json:"primarySubnetId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// VlanNumber - A VLAN's number as recorded within the SoftLayer network. This is configured directly
	// on SoftLayer's networking equipment and should not be confused with a VLAN's ''id'' property.
	VlanNumber int `json:"vlanNumber,omitempty"`

	// AccountId - The internal identifier of the SoftLayer customer account that a is associated with.
	AccountId int `json:"accountId,omitempty"`

	// TotalPrimaryIpAddressCount - no documentation
	TotalPrimaryIpAddressCount uint `json:"totalPrimaryIpAddressCount,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Network_Vlan_Type `json:"type,omitempty"`

	// VirtualGuests - All of the Virtual Servers that are connected to a
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests,omitempty"`

	// TagReferences - no documentation
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences,omitempty"`

	// SecondaryRouter - The secondary router that a is associated with. Every SoftLayer is connected to
	// more than one router for greater network redundancy.
	SecondaryRouter *SoftLayer_Hardware `json:"secondaryRouter,omitempty"`

	// FirewallRuleCount - A count of the currently running rule set of a firewalled
	FirewallRuleCount uint64 `json:"firewallRuleCount,omitempty"`

	// GuestNetworkComponentCount - A count of the networking components that are connected to a
	GuestNetworkComponentCount uint64 `json:"guestNetworkComponentCount,omitempty"`

	// PrimarySubnetVersion6 - A VLAN's primary IPv6 subnet. Some VLAN's may not have a primary IPv6
	// subnet.
	PrimarySubnetVersion6 *SoftLayer_Network_Subnet `json:"primarySubnetVersion6,omitempty"`

	// Subnets - no documentation
	Subnets []*SoftLayer_Network_Subnet `json:"subnets,omitempty"`

	// ProtectedIpAddressCount - no documentation
	ProtectedIpAddressCount uint64 `json:"protectedIpAddressCount,omitempty"`

	// ResourceGroupMemberCount - A count of the resource group member for a network vlan.
	ResourceGroupMemberCount uint64 `json:"resourceGroupMemberCount,omitempty"`

	// ProtectedIpAddresses - <nil>
	ProtectedIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"protectedIpAddresses,omitempty"`

	// PrimarySubnets - <nil>
	PrimarySubnets []*SoftLayer_Network_Subnet `json:"primarySubnets,omitempty"`

	// FirewallGuestNetworkComponents - no documentation
	FirewallGuestNetworkComponents []*SoftLayer_Network_Component_Firewall `json:"firewallGuestNetworkComponents,omitempty"`

	// NetworkComponentTrunks - The network components that are connected to this through a trunk.
	NetworkComponentTrunks []*SoftLayer_Network_Component_Network_Vlan_Trunk `json:"networkComponentTrunks,omitempty"`

	// PrivateNetworkGateways - no documentation
	PrivateNetworkGateways []*SoftLayer_Network_Gateway `json:"privateNetworkGateways,omitempty"`

	// ScaleVlans - no documentation
	ScaleVlans []*SoftLayer_Scale_Network_Vlan `json:"scaleVlans,omitempty"`

	// FirewallRules - no documentation
	FirewallRules []*SoftLayer_Network_Vlan_Firewall_Rule `json:"firewallRules,omitempty"`

	// HighAvailabilityFirewallFlag - <nil>
	HighAvailabilityFirewallFlag bool `json:"highAvailabilityFirewallFlag,omitempty"`

	// FirewallGuestNetworkComponentCount - no documentation
	FirewallGuestNetworkComponentCount uint64 `json:"firewallGuestNetworkComponentCount,omitempty"`

	// ResourceGroupCount - A count of the resource groups in which this is a member.
	ResourceGroupCount uint64 `json:"resourceGroupCount,omitempty"`

	// SubnetCount - A count of all of the subnets that exist as interfaces.
	SubnetCount uint64 `json:"subnetCount,omitempty"`

	// GuestNetworkComponents - no documentation
	GuestNetworkComponents []*SoftLayer_Virtual_Guest_Network_Component `json:"guestNetworkComponents,omitempty"`

	// FirewallNetworkComponents - no documentation
	FirewallNetworkComponents []*SoftLayer_Network_Component_Firewall `json:"firewallNetworkComponents,omitempty"`

	// SanStorageCapabilityFlag - A flag indicating that a vlan can be assigned to a host that has SAN disk
	// functionality.
	SanStorageCapabilityFlag bool `json:"sanStorageCapabilityFlag,omitempty"`

	// FirewallNetworkComponentCount - no documentation
	FirewallNetworkComponentCount uint64 `json:"firewallNetworkComponentCount,omitempty"`

	// ExtensionRouter - no documentation
	ExtensionRouter *SoftLayer_Hardware_Router `json:"extensionRouter,omitempty"`

	// PrimaryRouter - The primary router that a is associated with. Every SoftLayer is connected to more
	// than one router for greater network redundancy.
	PrimaryRouter *SoftLayer_Hardware_Router `json:"primaryRouter,omitempty"`

	// NetworkComponentTrunkCount - A count of the network components that are connected to this through a
	// trunk.
	NetworkComponentTrunkCount uint64 `json:"networkComponentTrunkCount,omitempty"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount,omitempty"`

	// NetworkVlanFirewall - The Hardware Firewall (Dedicated) for a network vlan.
	NetworkVlanFirewall *SoftLayer_Network_Vlan_Firewall `json:"networkVlanFirewall,omitempty"`

	// SecondarySubnets - The subnets that exist as secondary interfaces on a
	SecondarySubnets []*SoftLayer_Network_Subnet `json:"secondarySubnets,omitempty"`

	// LocalDiskStorageCapabilityFlag - A flag indicating that a vlan can be assigned to a host that has
	// local disk functionality.
	LocalDiskStorageCapabilityFlag bool `json:"localDiskStorageCapabilityFlag,omitempty"`

	// NetworkSpace - Identifier to denote whether a is used for public or private connectivity.
	NetworkSpace string `json:"networkSpace,omitempty"`

	// PrimarySubnet - A VLAN's primary subnet. Each has at least one subnet, usually the subnet that is
	// assigned to a server or new IP address block when it's purchased.
	PrimarySubnet *SoftLayer_Network_Subnet `json:"primarySubnet,omitempty"`

	// AdditionalPrimarySubnetCount - A count of a VLAN's additional primary subnets. These are used to
	// extend the number of servers attached to the by adding more ip addresses to the primary IP address
	// pool.
	AdditionalPrimarySubnetCount uint64 `json:"additionalPrimarySubnetCount,omitempty"`

	// HardwareCount - A count of all of the hardware that exists on a Hardware is associated with a by its
	// networking components.
	HardwareCount uint64 `json:"hardwareCount,omitempty"`

	// PrivateNetworkGatewayCount - no documentation
	PrivateNetworkGatewayCount uint64 `json:"privateNetworkGatewayCount,omitempty"`

	// PublicNetworkGatewayCount - no documentation
	PublicNetworkGatewayCount uint64 `json:"publicNetworkGatewayCount,omitempty"`

	// FirewallInterfaces - no documentation
	FirewallInterfaces []*SoftLayer_Network_Firewall_Module_Context_Interface `json:"firewallInterfaces,omitempty"`

	// ScaleVlanCount - A count of collection of scale VLANs this applies to.
	ScaleVlanCount uint64 `json:"scaleVlanCount,omitempty"`

	// AttachedNetworkGatewayFlag - no documentation
	AttachedNetworkGatewayFlag bool `json:"attachedNetworkGatewayFlag,omitempty"`

	// Hardware - All of the hardware that exists on a Hardware is associated with a by its networking
	// components.
	Hardware []*SoftLayer_Hardware `json:"hardware,omitempty"`

	// ResourceGroupMember - no documentation
	ResourceGroupMember []*SoftLayer_Resource_Group_Member `json:"resourceGroupMember,omitempty"`

	// ResourceGroups - no documentation
	ResourceGroups []*SoftLayer_Resource_Group `json:"resourceGroups,omitempty"`

	// AttachedNetworkGateway - no documentation
	AttachedNetworkGateway *SoftLayer_Network_Gateway `json:"attachedNetworkGateway,omitempty"`

	// FirewallInterfaceCount - A count of a firewalled vlan's inbound/outbound interfaces.
	FirewallInterfaceCount uint64 `json:"firewallInterfaceCount,omitempty"`

	// PublicNetworkGateways - no documentation
	PublicNetworkGateways []*SoftLayer_Network_Gateway `json:"publicNetworkGateways,omitempty"`

	// NetworkComponents - no documentation
	NetworkComponents []*SoftLayer_Network_Component `json:"networkComponents,omitempty"`

	// NetworkComponentCount - A count of the networking components that are connected to a
	NetworkComponentCount uint64 `json:"networkComponentCount,omitempty"`

	// DedicatedFirewallFlag - A flag indicating that a network vlan is on a Hardware Firewall (Dedicated).
	DedicatedFirewallFlag int `json:"dedicatedFirewallFlag,omitempty"`

	// SecondarySubnetCount - A count of the subnets that exist as secondary interfaces on a
	SecondarySubnetCount uint64 `json:"secondarySubnetCount,omitempty"`

	// PrimarySubnetCount - no documentation
	PrimarySubnetCount uint64 `json:"primarySubnetCount,omitempty"`

	// AttachedNetworkGatewayVlan - The inside record if this is inside a network gateway.
	AttachedNetworkGatewayVlan *SoftLayer_Network_Gateway_Vlan `json:"attachedNetworkGatewayVlan,omitempty"`

	// VirtualGuestCount - A count of all of the Virtual Servers that are connected to a
	VirtualGuestCount uint64 `json:"virtualGuestCount,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// AdditionalPrimarySubnets - A VLAN's additional primary subnets. These are used to extend the number
	// of servers attached to the by adding more ip addresses to the primary IP address pool.
	AdditionalPrimarySubnets []*SoftLayer_Network_Subnet `json:"additionalPrimarySubnets,omitempty"`
}

func (softlayer_network_vlan *SoftLayer_Network_Vlan) String() string {
	return "SoftLayer_Network_Vlan"
}
