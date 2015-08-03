package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Subnet - The SoftLayer_Network_Subnet data type contains general information
// relating to a single SoftLayer subnet. Personal information in this type such as names, addresses,
// and phone numbers are assigned to the account only and not to users belonging to the account.
type SoftLayer_Network_Subnet struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// ActiveRegistration - If present, the active registration for this subnet.
	ActiveRegistration *SoftLayer_Network_Subnet_Registration `json:"activeRegistration"`

	// ActiveSwipTransaction - All the swip transactions associated with a subnet that are still active.
	ActiveSwipTransaction *SoftLayer_Network_Subnet_Swip_Transaction `json:"activeSwipTransaction"`

	// ActiveTransaction - no documentation
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction"`

	// AddressSpace - Identifier which distinguishes whether the subnet is public or private address space.
	AddressSpace string `json:"addressSpace"`

	// AllowedHost - The SoftLayer_Network_Storage_Allowed_Host information to connect this Subnet to
	// Network Storage supporting access control lists.
	AllowedHost *SoftLayer_Network_Storage_Allowed_Host `json:"allowedHost"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// BoundDescendantCount - no documentation
	BoundDescendantCount uint64 `json:"boundDescendantCount"`

	// BoundDescendants - <nil>
	BoundDescendants []*SoftLayer_Network_Subnet `json:"boundDescendants"`

	// BoundRouterCount - no documentation
	BoundRouterCount uint64 `json:"boundRouterCount"`

	// BoundRouterFlag - Whether or not this subnet is associated with a router. Subnets that are not
	// associated with a router cannot be routed.
	BoundRouterFlag bool `json:"boundRouterFlag"`

	// BoundRouters - <nil>
	BoundRouters []*SoftLayer_Hardware `json:"boundRouters"`

	// BroadcastAddress - The last IP address in a subnet is the subnet's broadcast address. This is an IP
	// address that will broadcast network requests to the entire subnet and may not be assigned to a
	// network interface.
	BroadcastAddress string `json:"broadcastAddress"`

	// Children - <nil>
	Children []*SoftLayer_Network_Subnet `json:"children"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount"`

	// Cidr - A subnet's Classless Inter-Domain Routing prefix. This is a number between 0 and 32
	// signifying the number of bits in a subnet's netmask. These bits separate a subnet's network address
	// from it's host addresses. It performs the same function as the ''netmask'' property, but is
	// represented as an integer.
	Cidr int `json:"cidr"`

	// Datacenter - no documentation
	Datacenter *SoftLayer_Location_Datacenter `json:"datacenter"`

	// DescendantCount - no documentation
	DescendantCount uint64 `json:"descendantCount"`

	// Descendants - <nil>
	Descendants []*SoftLayer_Network_Subnet `json:"descendants"`

	// DisplayLabel - <nil>
	DisplayLabel string `json:"displayLabel"`

	// EndPointIpAddress - no documentation
	EndPointIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"endPointIpAddress"`

	// Gateway - A subnet's gateway address. This is an IP address that belongs to the router on the subnet
	// and may not be assigned to a network interface.
	Gateway string `json:"gateway"`

	// GlobalIpRecord - <nil>
	GlobalIpRecord *SoftLayer_Network_Subnet_IpAddress_Global `json:"globalIpRecord"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// HardwareCount - A count of the hardware that this subnet is routed to.
	HardwareCount uint64 `json:"hardwareCount"`

	// Id - no documentation
	Id int `json:"id"`

	// IpAddressCount - A count of all the ip addresses associated with a subnet.
	IpAddressCount uint64 `json:"ipAddressCount"`

	// IpAddresses - no documentation
	IpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"ipAddresses"`

	// IsCustomerOwned - <nil>
	IsCustomerOwned bool `json:"isCustomerOwned"`

	// IsCustomerRoutable - <nil>
	IsCustomerRoutable bool `json:"isCustomerRoutable"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Netmask - A bitmask in dotted-quad format that is used to separate a subnet's network address from
	// it's host addresses. This performs the same function as the ''cidr'' property, but is expressed in a
	// string format.
	Netmask string `json:"netmask"`

	// NetworkComponent - no documentation
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`

	// NetworkComponentFirewall - no documentation
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall"`

	// NetworkIdentifier - A subnet's network identifier. This is the first IP address of a subnet and may
	// not be assigned to a network interface.
	NetworkIdentifier string `json:"networkIdentifier"`

	// NetworkProtectionAddressCount - no documentation
	NetworkProtectionAddressCount uint64 `json:"networkProtectionAddressCount"`

	// NetworkProtectionAddresses - <nil>
	NetworkProtectionAddresses []*SoftLayer_Network_Protection_Address `json:"networkProtectionAddresses"`

	// NetworkTunnelContextCount - A count of iPSec network tunnels that have access to a private subnet.
	NetworkTunnelContextCount uint64 `json:"networkTunnelContextCount"`

	// NetworkTunnelContexts - IPSec network tunnels that have access to a private subnet.
	NetworkTunnelContexts []*SoftLayer_Network_Tunnel_Module_Context `json:"networkTunnelContexts"`

	// NetworkVlan - no documentation
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// NetworkVlanId - no documentation
	NetworkVlanId int `json:"networkVlanId"`

	// Note - no documentation
	Note string `json:"note"`

	// ProtectedIpAddressCount - no documentation
	ProtectedIpAddressCount uint64 `json:"protectedIpAddressCount"`

	// ProtectedIpAddresses - <nil>
	ProtectedIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"protectedIpAddresses"`

	// RegionalInternetRegistry - <nil>
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry"`

	// RegistrationCount - A count of all registrations that have been created for this subnet.
	RegistrationCount uint64 `json:"registrationCount"`

	// Registrations - All registrations that have been created for this subnet.
	Registrations []*SoftLayer_Network_Subnet_Registration `json:"registrations"`

	// ResourceGroupCount - A count of the resource groups in which this subnet is a member.
	ResourceGroupCount uint64 `json:"resourceGroupCount"`

	// ResourceGroups - The resource groups in which this subnet is a member.
	ResourceGroups []*SoftLayer_Resource_Group `json:"resourceGroups"`

	// ReverseDomain - The reverse DNS domain associated with this subnet.
	ReverseDomain *SoftLayer_Dns_Domain `json:"reverseDomain"`

	// RoleKeyName - An identifier of the role the subnet is within. Roles dictate how a subnet may be
	// used.
	RoleKeyName string `json:"roleKeyName"`

	// RoleName - The name of the role the subnet is within. Roles dictate how a subnet may be used.
	RoleName string `json:"roleName"`

	// RoutingTypeKeyName - The identifier for the type of route then subnet is currently configured for.
	RoutingTypeKeyName string `json:"routingTypeKeyName"`

	// RoutingTypeName - The name for the type of route then subnet is currently configured for.
	RoutingTypeName string `json:"routingTypeName"`

	// SortOrder - A subnet can be one of several types. and The type determines the order in which many
	// subnets are sorted in the SoftLayer customer portal. This groups subnets of similar type together.
	SortOrder string `json:"sortOrder"`

	// SubnetType - A subnet can be one of several types. and A subnet is the primary network bound to a
	// within the softlayer network. An subnet is bound to a network to augment the pool of available
	// primary IP addresses that may be assigned to a server. A subnet is any of the secondary subnet's
	// bound to a interface. A subnet is a portable subnet that can be routed to any server on a vlan. A
	// subnet also doesn't exist as a interface, but is routed directly to a instead of a single IP address
	// by SoftLayer's routers.
	SubnetType string `json:"subnetType"`

	// SwipTransaction - All the swip transactions associated with a subnet.
	SwipTransaction []*SoftLayer_Network_Subnet_Swip_Transaction `json:"swipTransaction"`

	// SwipTransactionCount - A count of all the swip transactions associated with a subnet.
	SwipTransactionCount uint64 `json:"swipTransactionCount"`

	// TotalIpAddresses - The number of IP addresses contained within this subnet.
	TotalIpAddresses uint `json:"totalIpAddresses"`

	// UnboundDescendantCount - no documentation
	UnboundDescendantCount uint64 `json:"unboundDescendantCount"`

	// UnboundDescendants - <nil>
	UnboundDescendants []*SoftLayer_Network_Subnet `json:"unboundDescendants"`

	// UsableIpAddressCount - The number of IP addresses that can be addressed within this subnet.
	UsableIpAddressCount uint `json:"usableIpAddressCount"`

	// Version - This is the Internet Protocol version. Current values may be either 4 or 6.
	Version int `json:"version"`

	// VirtualGuestCount - A count of the Virtual Servers that this subnet is routed to.
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`
}

func (softlayer_network_subnet *SoftLayer_Network_Subnet) String() string {
	return "SoftLayer_Network_Subnet"
}

// CreateReverseDomainRecords - no documentation
func (softlayer_network_subnet *SoftLayer_Network_Subnet) CreateReverseDomainRecords(ctx *slapi.RequestContext) (*SoftLayer_Dns_Domain_Reverse, error) {
	var returnValue *SoftLayer_Dns_Domain_Reverse
	return returnValue, nil
}

// CreateSubnetRouteUpdateTransaction - This function is used to create a new transaction to modify a
// subnet route. Routes are updated in one to two minutes depending on the number of transactions that
// are pending for a router. Usage of this function is restricted and may only be called from
// authorized accounts. It is not available for general API users without justification and consent
// from a SoftLayer representative.
func (softlayer_network_subnet *SoftLayer_Network_Subnet) CreateSubnetRouteUpdateTransaction(ctx *slapi.RequestContext, newEndPointIpAddress string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateSwipTransaction - This function is used to create a new SoftLayer transaction to register your
// data with transactions can only be initiated on subnets that contain more than 8 IP addresses.
func (softlayer_network_subnet *SoftLayer_Network_Subnet) CreateSwipTransaction(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditNote - no documentation
func (softlayer_network_subnet *SoftLayer_Network_Subnet) EditNote(ctx *slapi.RequestContext, note string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// FindAllSubnetsAndActiveSwipTransactionStatus - Retrieve a list of a SoftLayer customer's subnets
// along with their transaction statuses. This is a shortcut method that combines the
// SoftLayer_Network_Subnet retrieval methods along with [[object masks]] to retrieve their subnets'
// associated transactions as well. This is a special function built for SoftLayer's use on the section
// of the customer portal, but may also be useful for API users looking for the same data.
func (softlayer_network_subnet *SoftLayer_Network_Subnet) FindAllSubnetsAndActiveSwipTransactionStatus(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Subnet, error) {
	var returnValue []*SoftLayer_Network_Subnet
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Subnet object whose ID number corresponds to
// the ID number of the init parameter passed to the SoftLayer_Network_Subnet service. You can only
// retrieve the subnet whose vlan is associated with the account that you portal user is assigned to.
func (softlayer_network_subnet *SoftLayer_Network_Subnet) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Subnet, error) {
	var returnValue *SoftLayer_Network_Subnet
	return returnValue, nil
}

// GetReverseDomainRecords - Retrieve all reverse DNS records associated with a subnet.
func (softlayer_network_subnet *SoftLayer_Network_Subnet) GetReverseDomainRecords(ctx *slapi.RequestContext) ([]*SoftLayer_Dns_Domain, error) {
	var returnValue []*SoftLayer_Dns_Domain
	return returnValue, nil
}

// GetRoutableEndpointIpAddresses - getRoutableEndpointAddresses retrieves valid routable endpoint
// addresses for a subnet. You may use any IP address in a portable subnet, but may not use the network
// identifier, gateway, or broadcast address for primary and secondary on subnets.
func (softlayer_network_subnet *SoftLayer_Network_Subnet) GetRoutableEndpointIpAddresses(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Subnet_IpAddress, error) {
	var returnValue []*SoftLayer_Network_Subnet_IpAddress
	return returnValue, nil
}

// GetSubnetForIpAddress - Retrieve the subnet associated with an IP address. You may only retrieve
// subnets assigned to your SoftLayer customer account.
func (softlayer_network_subnet *SoftLayer_Network_Subnet) GetSubnetForIpAddress(ctx *slapi.RequestContext, ipAddress string) (*SoftLayer_Network_Subnet, error) {
	var returnValue *SoftLayer_Network_Subnet
	return returnValue, nil
}
