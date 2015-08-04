package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Subnet - The SoftLayer_Network_Subnet data type contains general information
// relating to a single SoftLayer subnet. Personal information in this type such as names, addresses,
// and phone numbers are assigned to the account only and not to users belonging to the account.
type SoftLayer_Network_Subnet struct {

	// TotalIpAddresses - The number of IP addresses contained within this subnet.
	TotalIpAddresses uint `json:"totalIpAddresses,omitempty"`

	// Gateway - A subnet's gateway address. This is an IP address that belongs to the router on the subnet
	// and may not be assigned to a network interface.
	Gateway string `json:"gateway,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// SubnetType - A subnet can be one of several types. and A subnet is the primary network bound to a
	// within the softlayer network. An subnet is bound to a network to augment the pool of available
	// primary IP addresses that may be assigned to a server. A subnet is any of the secondary subnet's
	// bound to a interface. A subnet is a portable subnet that can be routed to any server on a vlan. A
	// subnet also doesn't exist as a interface, but is routed directly to a instead of a single IP address
	// by SoftLayer's routers.
	SubnetType string `json:"subnetType,omitempty"`

	// Version - This is the Internet Protocol version. Current values may be either 4 or 6.
	Version int `json:"version,omitempty"`

	// IsCustomerOwned - <nil>
	IsCustomerOwned bool `json:"isCustomerOwned,omitempty"`

	// Note - no documentation
	Note string `json:"note,omitempty"`

	// BroadcastAddress - The last IP address in a subnet is the subnet's broadcast address. This is an IP
	// address that will broadcast network requests to the entire subnet and may not be assigned to a
	// network interface.
	BroadcastAddress string `json:"broadcastAddress,omitempty"`

	// Netmask - A bitmask in dotted-quad format that is used to separate a subnet's network address from
	// it's host addresses. This performs the same function as the ''cidr'' property, but is expressed in a
	// string format.
	Netmask string `json:"netmask,omitempty"`

	// Cidr - A subnet's Classless Inter-Domain Routing prefix. This is a number between 0 and 32
	// signifying the number of bits in a subnet's netmask. These bits separate a subnet's network address
	// from it's host addresses. It performs the same function as the ''netmask'' property, but is
	// represented as an integer.
	Cidr int `json:"cidr,omitempty"`

	// IsCustomerRoutable - <nil>
	IsCustomerRoutable bool `json:"isCustomerRoutable,omitempty"`

	// NetworkIdentifier - A subnet's network identifier. This is the first IP address of a subnet and may
	// not be assigned to a network interface.
	NetworkIdentifier string `json:"networkIdentifier,omitempty"`

	// NetworkVlanId - no documentation
	NetworkVlanId int `json:"networkVlanId,omitempty"`

	// SortOrder - A subnet can be one of several types. and The type determines the order in which many
	// subnets are sorted in the SoftLayer customer portal. This groups subnets of similar type together.
	SortOrder string `json:"sortOrder,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// UsableIpAddressCount - The number of IP addresses that can be addressed within this subnet.
	UsableIpAddressCount uint `json:"usableIpAddressCount,omitempty"`
}

func (softlayer_network_subnet *SoftLayer_Network_Subnet) String() string {
	return "SoftLayer_Network_Subnet"
}

// SoftLayer_Network_Subnet_Extended is SoftLayer_Network_Subnet with all maskable types.
type SoftLayer_Network_Subnet_Extended struct {
	SoftLayer_Network_Subnet

	// UnboundDescendants - <nil>
	UnboundDescendants []*SoftLayer_Network_Subnet `json:"unboundDescendants,omitempty"`

	// RegistrationCount - A count of all registrations that have been created for this subnet.
	RegistrationCount uint64 `json:"registrationCount,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// BoundRouters - <nil>
	BoundRouters []*SoftLayer_Hardware `json:"boundRouters,omitempty"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware,omitempty"`

	// NetworkTunnelContexts - IPSec network tunnels that have access to a private subnet.
	NetworkTunnelContexts []*SoftLayer_Network_Tunnel_Module_Context `json:"networkTunnelContexts,omitempty"`

	// BoundRouterFlag - Whether or not this subnet is associated with a router. Subnets that are not
	// associated with a router cannot be routed.
	BoundRouterFlag bool `json:"boundRouterFlag,omitempty"`

	// Children - <nil>
	Children []*SoftLayer_Network_Subnet `json:"children,omitempty"`

	// VirtualGuestCount - A count of the Virtual Servers that this subnet is routed to.
	VirtualGuestCount uint64 `json:"virtualGuestCount,omitempty"`

	// NetworkProtectionAddresses - <nil>
	NetworkProtectionAddresses []*SoftLayer_Network_Protection_Address `json:"networkProtectionAddresses,omitempty"`

	// NetworkVlan - no documentation
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan,omitempty"`

	// RegionalInternetRegistry - <nil>
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry,omitempty"`

	// EndPointIpAddress - no documentation
	EndPointIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"endPointIpAddress,omitempty"`

	// GlobalIpRecord - <nil>
	GlobalIpRecord *SoftLayer_Network_Subnet_IpAddress_Global `json:"globalIpRecord,omitempty"`

	// BoundDescendantCount - no documentation
	BoundDescendantCount uint64 `json:"boundDescendantCount,omitempty"`

	// ReverseDomain - The reverse DNS domain associated with this subnet.
	ReverseDomain *SoftLayer_Dns_Domain `json:"reverseDomain,omitempty"`

	// Registrations - All registrations that have been created for this subnet.
	Registrations []*SoftLayer_Network_Subnet_Registration `json:"registrations,omitempty"`

	// RoutingTypeKeyName - The identifier for the type of route then subnet is currently configured for.
	RoutingTypeKeyName string `json:"routingTypeKeyName,omitempty"`

	// ResourceGroupCount - A count of the resource groups in which this subnet is a member.
	ResourceGroupCount uint64 `json:"resourceGroupCount,omitempty"`

	// UnboundDescendantCount - no documentation
	UnboundDescendantCount uint64 `json:"unboundDescendantCount,omitempty"`

	// AllowedHost - The SoftLayer_Network_Storage_Allowed_Host information to connect this Subnet to
	// Network Storage supporting access control lists.
	AllowedHost *SoftLayer_Network_Storage_Allowed_Host `json:"allowedHost,omitempty"`

	// Descendants - <nil>
	Descendants []*SoftLayer_Network_Subnet `json:"descendants,omitempty"`

	// BoundRouterCount - no documentation
	BoundRouterCount uint64 `json:"boundRouterCount,omitempty"`

	// ActiveTransaction - no documentation
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction,omitempty"`

	// ProtectedIpAddresses - <nil>
	ProtectedIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"protectedIpAddresses,omitempty"`

	// DescendantCount - no documentation
	DescendantCount uint64 `json:"descendantCount,omitempty"`

	// DisplayLabel - <nil>
	DisplayLabel string `json:"displayLabel,omitempty"`

	// NetworkComponent - no documentation
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent,omitempty"`

	// HardwareCount - A count of the hardware that this subnet is routed to.
	HardwareCount uint64 `json:"hardwareCount,omitempty"`

	// NetworkTunnelContextCount - A count of iPSec network tunnels that have access to a private subnet.
	NetworkTunnelContextCount uint64 `json:"networkTunnelContextCount,omitempty"`

	// RoleKeyName - An identifier of the role the subnet is within. Roles dictate how a subnet may be
	// used.
	RoleKeyName string `json:"roleKeyName,omitempty"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests,omitempty"`

	// IpAddressCount - A count of all the ip addresses associated with a subnet.
	IpAddressCount uint64 `json:"ipAddressCount,omitempty"`

	// ActiveSwipTransaction - All the swip transactions associated with a subnet that are still active.
	ActiveSwipTransaction *SoftLayer_Network_Subnet_Swip_Transaction `json:"activeSwipTransaction,omitempty"`

	// AddressSpace - Identifier which distinguishes whether the subnet is public or private address space.
	AddressSpace string `json:"addressSpace,omitempty"`

	// NetworkProtectionAddressCount - no documentation
	NetworkProtectionAddressCount uint64 `json:"networkProtectionAddressCount,omitempty"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// RoleName - The name of the role the subnet is within. Roles dictate how a subnet may be used.
	RoleName string `json:"roleName,omitempty"`

	// ProtectedIpAddressCount - no documentation
	ProtectedIpAddressCount uint64 `json:"protectedIpAddressCount,omitempty"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location_Datacenter `json:"datacenter,omitempty"`

	// NetworkComponentFirewall - no documentation
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall,omitempty"`

	// SwipTransaction - All the swip transactions associated with a subnet.
	SwipTransaction []*SoftLayer_Network_Subnet_Swip_Transaction `json:"swipTransaction,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// ResourceGroups - The resource groups in which this subnet is a member.
	ResourceGroups []*SoftLayer_Resource_Group `json:"resourceGroups,omitempty"`

	// IpAddresses - no documentation
	IpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"ipAddresses,omitempty"`

	// SwipTransactionCount - A count of all the swip transactions associated with a subnet.
	SwipTransactionCount uint64 `json:"swipTransactionCount,omitempty"`

	// ActiveRegistration - If present, the active registration for this subnet.
	ActiveRegistration *SoftLayer_Network_Subnet_Registration `json:"activeRegistration,omitempty"`

	// BoundDescendants - <nil>
	BoundDescendants []*SoftLayer_Network_Subnet `json:"boundDescendants,omitempty"`

	// RoutingTypeName - The name for the type of route then subnet is currently configured for.
	RoutingTypeName string `json:"routingTypeName,omitempty"`
}

func (softlayer_network_subnet *SoftLayer_Network_Subnet_Extended) String() string {
	return "SoftLayer_Network_Subnet"
}
