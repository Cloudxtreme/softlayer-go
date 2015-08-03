package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Gateway - <nil>
type SoftLayer_Network_Gateway struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The internal identifier of the account assigned to this gateway.
	AccountId int `json:"accountId"`

	// GroupNumber - The group number for this gateway. This is set internally and cannot be provided on
	// create.
	GroupNumber int `json:"groupNumber"`

	// Id - no documentation
	Id int `json:"id"`

	// InsideVlanCount - no documentation
	InsideVlanCount uint64 `json:"insideVlanCount"`

	// InsideVlans - no documentation
	InsideVlans []*SoftLayer_Network_Gateway_Vlan `json:"insideVlans"`

	// MemberCount - no documentation
	MemberCount uint64 `json:"memberCount"`

	// Members - no documentation
	Members []*SoftLayer_Network_Gateway_Member `json:"members"`

	// Name - A gateway's name. This is required on create and can be no more than 255 characters.
	Name string `json:"name"`

	// NetworkSpace - A gateway's network space. Currently, only 'private' or 'both' is allowed. When this
	// value is 'private', it is a backend gateway only. Otherwise, it is a gateway for both frontend and
	// backend traffic.
	NetworkSpace string `json:"networkSpace"`

	// PrivateIpAddress - no documentation
	PrivateIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"privateIpAddress"`

	// PrivateIpAddressId - The internal identifier of the private IP address for this gateway.
	PrivateIpAddressId int `json:"privateIpAddressId"`

	// PrivateVlan - no documentation
	PrivateVlan *SoftLayer_Network_Vlan `json:"privateVlan"`

	// PrivateVlanId - The internal identifier of the private for this gateway.
	PrivateVlanId int `json:"privateVlanId"`

	// PublicIpAddress - no documentation
	PublicIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"publicIpAddress"`

	// PublicIpAddressId - The internal identifier of the public IP address for this gateway.
	PublicIpAddressId int `json:"publicIpAddressId"`

	// PublicIpv6Address - no documentation
	PublicIpv6Address *SoftLayer_Network_Subnet_IpAddress `json:"publicIpv6Address"`

	// PublicIpv6AddressId - The internal identifier of the public IPv6 address for this gateway.
	PublicIpv6AddressId int `json:"publicIpv6AddressId"`

	// PublicVlan - no documentation
	PublicVlan *SoftLayer_Network_Vlan `json:"publicVlan"`

	// PublicVlanId - The internal identifier of the public for this gateway. This is set internally and
	// cannot be provided on create.
	PublicVlanId int `json:"publicVlanId"`

	// Status - no documentation
	Status *SoftLayer_Network_Gateway_Status `json:"status"`

	// StatusId - The current status of this gateway. This is always active unless there is a process
	// running to change the gateway. This can not be set on creation.
	StatusId int `json:"statusId"`
}

func (softlayer_network_gateway *SoftLayer_Network_Gateway) String() string {
	return "SoftLayer_Network_Gateway"
}

// BypassAllVlans - Start the asynchronous process to bypass all VLANs. Any VLANs that are already
// bypassed will be ignored. The status field can be checked for progress.
func (softlayer_network_gateway *SoftLayer_Network_Gateway) BypassAllVlans(ctx *slapi.RequestContext) error {
	return nil
}

// BypassVlans - Start the asynchronous process to bypass the provided VLANs. The VLANs must already be
// attached. Any VLANs that are already bypassed will be ignored. The status field can be checked for
// progress.
func (softlayer_network_gateway *SoftLayer_Network_Gateway) BypassVlans(ctx *slapi.RequestContext, vlans []SoftLayer_Network_Gateway_Vlan) error {
	return nil
}

// CreateObject - Create and return a new gateway. This object can be created with any number of
// members or VLANs, but they all must be in the same pod. By creating a gateway with members and/or
// VLANs attached, it is the equivalent of individually calling their createObject methods except this
// will start a single asynchronous process to setup the gateway. The status of this process can be
// checked using the status field.
func (softlayer_network_gateway *SoftLayer_Network_Gateway) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Gateway) (*SoftLayer_Network_Gateway, error) {
	var returnValue *SoftLayer_Network_Gateway
	return returnValue, nil
}

// EditObject - Edit this gateway. Currently, the only value that can be edited is the name.
func (softlayer_network_gateway *SoftLayer_Network_Gateway) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Gateway) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_gateway *SoftLayer_Network_Gateway) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Gateway, error) {
	var returnValue *SoftLayer_Network_Gateway
	return returnValue, nil
}

// GetPossibleInsideVlans - Get all VLANs that can become inside VLANs on this gateway. This means the
// must not already be an inside on the same router as this gateway, not a gateway transit and not
// firewalled.
func (softlayer_network_gateway *SoftLayer_Network_Gateway) GetPossibleInsideVlans(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Vlan, error) {
	var returnValue []*SoftLayer_Network_Vlan
	return returnValue, nil
}

// UnbypassAllVlans - Start the asynchronous process to unbypass all VLANs. Any VLANs that are already
// unbypassed will be ignored. The status field can be checked for progress.
func (softlayer_network_gateway *SoftLayer_Network_Gateway) UnbypassAllVlans(ctx *slapi.RequestContext) error {
	return nil
}

// UnbypassVlans - Start the asynchronous process to unbypass the provided VLANs. The VLANs must
// already be attached. Any VLANs that are already unbypassed will be ignored. The status field can be
// checked for progress.
func (softlayer_network_gateway *SoftLayer_Network_Gateway) UnbypassVlans(ctx *slapi.RequestContext, vlans []SoftLayer_Network_Gateway_Vlan) error {
	return nil
}
