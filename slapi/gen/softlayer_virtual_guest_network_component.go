package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Guest_Network_Component - The virtual guest network component data type presents
// the structure in which all computing instance network components are presented. Internally, the
// structure supports various virtualization platforms with no change to external interaction. A guest,
// also known as a virtual server, represents an allocation of resources on a virtual host.
type SoftLayer_Virtual_Guest_Network_Component struct {

	// CreateDate - The date a computing instance's network component was created.
	CreateDate *time.Time `json:"createDate"`

	// Guest - The computing instance that this network component exists on.
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// GuestId - The unique ID of the [[SoftLayer_Virtual_Guest|computing instance]] that this network
	// component belongs to.
	GuestId int `json:"guestId"`

	// HighAvailabilityFirewallFlag - <nil>
	HighAvailabilityFirewallFlag bool `json:"highAvailabilityFirewallFlag"`

	// Id - A computing instance's network component's unique ID.
	Id int `json:"id"`

	// IpAddressBindingCount - A count of the records of all IP addresses bound to a computing instance's
	// network component.
	IpAddressBindingCount uint64 `json:"ipAddressBindingCount"`

	// IpAddressBindings - The records of all IP addresses bound to a computing instance's network
	// component.
	IpAddressBindings []*SoftLayer_Virtual_Guest_Network_Component_IpAddress `json:"ipAddressBindings"`

	// MacAddress - A computing instance network component's unique MAC address.
	MacAddress string `json:"macAddress"`

	// MaxSpeed - A computing instance network component's maximum allowed speed, measured in Mbit per
	// second. ''maxSpeed'' is determined by the capabilities of the network interface and the port speed
	// purchased on your SoftLayer computing instance.
	MaxSpeed int `json:"maxSpeed"`

	// ModifyDate - The date a computing instance's network component was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - A computing instance network component's short name. This is usually ''eth''. Use this in
	// conjunction with the ''port'' property to identify a network component. For instance, the "eth0"
	// interface on a server has the network component name "eth" and port 0.
	Name string `json:"name"`

	// NetworkComponentFirewall - no documentation
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall"`

	// NetworkId - A computing instance's network component's [[SoftLayer_Virtual_Network|network]] ID
	NetworkId int `json:"networkId"`

	// NetworkVlan - The that a computing instance network component's subnet is associated with.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// Port - A computing instance network component's port number. Most computing instances have more than
	// one network interface. The port property separates these interfaces. Use this in conjunction with
	// the ''name'' property to identify a network component. For instance, the "eth0" interface on a
	// server has the network component name "eth" and port 0.
	Port int `json:"port"`

	// PrimaryIpAddress - A computing instance network component's primary IP address.
	PrimaryIpAddress string `json:"primaryIpAddress"`

	// PrimaryIpAddressRecord - <nil>
	PrimaryIpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryIpAddressRecord"`

	// PrimarySubnet - A network component's subnet for its primary IP address
	PrimarySubnet *SoftLayer_Network_Subnet `json:"primarySubnet"`

	// PrimaryVersion6IpAddressRecord - A network component's primary IPv6 IP address record.
	PrimaryVersion6IpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryVersion6IpAddressRecord"`

	// Router - no documentation
	Router *SoftLayer_Hardware_Router `json:"router"`

	// Speed - A computing instance network component's speed, measured in Mbit per second.
	Speed int `json:"speed"`

	// Status - A computing instance network component's status. This can take one of four possible values:
	// network components are enabled and in use on a cloud instance. status components have been
	// administratively disabled by SoftLayer accounting or abuse. components have been administratively
	// disabled by you, the user. You should never see a network interface in state. If you happen to see
	// one please contact SoftLayer support.
	Status string `json:"status"`

	// SubnetCount - A count of a network component's subnets. A subnet is a group of IP addresses
	SubnetCount uint64 `json:"subnetCount"`

	// Subnets - A network component's subnets. A subnet is a group of IP addresses
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// Uuid - A computing instance's network component's unique ID on a virtualization platform.
	Uuid string `json:"uuid"`
}

// Disable - Completely restrict all incoming and outgoing bandwidth traffic to a network component
func (softlayer_virtual_guest_network_component *SoftLayer_Virtual_Guest_Network_Component) Disable() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Enable - Allow incoming and outgoing bandwidth traffic to a network component
func (softlayer_virtual_guest_network_component *SoftLayer_Virtual_Guest_Network_Component) Enable() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_virtual_guest_network_component *SoftLayer_Virtual_Guest_Network_Component) GetObject() (*SoftLayer_Virtual_Guest_Network_Component, error) {
	var returnValue *SoftLayer_Virtual_Guest_Network_Component
	return returnValue, nil
}

// IsPingable - Issues a ping command and returns the success (true) or failure (false) of the ping
// command.
func (softlayer_virtual_guest_network_component *SoftLayer_Virtual_Guest_Network_Component) IsPingable() (bool, error) {
	var returnValue bool
	return returnValue, nil
}
