package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Guest_Network_Component - The virtual guest network component data type presents
// the structure in which all computing instance network components are presented. Internally, the
// structure supports various virtualization platforms with no change to external interaction. A guest,
// also known as a virtual server, represents an allocation of resources on a virtual host.
type SoftLayer_Virtual_Guest_Network_Component struct {

	// MacAddress - A computing instance network component's unique MAC address.
	MacAddress string `json:"macAddress,omitempty"`

	// Port - A computing instance network component's port number. Most computing instances have more than
	// one network interface. The port property separates these interfaces. Use this in conjunction with
	// the ''name'' property to identify a network component. For instance, the "eth0" interface on a
	// server has the network component name "eth" and port 0.
	Port int `json:"port,omitempty"`

	// Status - A computing instance network component's status. This can take one of four possible values:
	// network components are enabled and in use on a cloud instance. status components have been
	// administratively disabled by SoftLayer accounting or abuse. components have been administratively
	// disabled by you, the user. You should never see a network interface in state. If you happen to see
	// one please contact SoftLayer support.
	Status string `json:"status,omitempty"`

	// CreateDate - The date a computing instance's network component was created.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// GuestId - The unique ID of the [[SoftLayer_Virtual_Guest|computing instance]] that this network
	// component belongs to.
	GuestId int `json:"guestId,omitempty"`

	// Id - A computing instance's network component's unique ID.
	Id int `json:"id,omitempty"`

	// NetworkId - A computing instance's network component's [[SoftLayer_Virtual_Network|network]] ID
	NetworkId int `json:"networkId,omitempty"`

	// Speed - A computing instance network component's speed, measured in Mbit per second.
	Speed int `json:"speed,omitempty"`

	// Uuid - A computing instance's network component's unique ID on a virtualization platform.
	Uuid string `json:"uuid,omitempty"`

	// MaxSpeed - A computing instance network component's maximum allowed speed, measured in Mbit per
	// second. ''maxSpeed'' is determined by the capabilities of the network interface and the port speed
	// purchased on your SoftLayer computing instance.
	MaxSpeed int `json:"maxSpeed,omitempty"`

	// ModifyDate - The date a computing instance's network component was last modified.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - A computing instance network component's short name. This is usually ''eth''. Use this in
	// conjunction with the ''port'' property to identify a network component. For instance, the "eth0"
	// interface on a server has the network component name "eth" and port 0.
	Name string `json:"name,omitempty"`
}

func (softlayer_virtual_guest_network_component *SoftLayer_Virtual_Guest_Network_Component) String() string {
	return "SoftLayer_Virtual_Guest_Network_Component"
}

// SoftLayer_Virtual_Guest_Network_Component_Extended is SoftLayer_Virtual_Guest_Network_Component with all maskable types.
type SoftLayer_Virtual_Guest_Network_Component_Extended struct {
	SoftLayer_Virtual_Guest_Network_Component

	// SubnetCount - A count of a network component's subnets. A subnet is a group of IP addresses
	SubnetCount uint64 `json:"subnetCount,omitempty"`

	// NetworkVlan - The that a computing instance network component's subnet is associated with.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan,omitempty"`

	// Subnets - A network component's subnets. A subnet is a group of IP addresses
	Subnets []*SoftLayer_Network_Subnet `json:"subnets,omitempty"`

	// IpAddressBindingCount - A count of the records of all IP addresses bound to a computing instance's
	// network component.
	IpAddressBindingCount uint64 `json:"ipAddressBindingCount,omitempty"`

	// PrimaryIpAddress - A computing instance network component's primary IP address.
	PrimaryIpAddress string `json:"primaryIpAddress,omitempty"`

	// PrimaryVersion6IpAddressRecord - A network component's primary IPv6 IP address record.
	PrimaryVersion6IpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryVersion6IpAddressRecord,omitempty"`

	// Router - no documentation
	Router *SoftLayer_Hardware_Router `json:"router,omitempty"`

	// NetworkComponentFirewall - no documentation
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall,omitempty"`

	// PrimaryIpAddressRecord - <nil>
	PrimaryIpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryIpAddressRecord,omitempty"`

	// PrimarySubnet - A network component's subnet for its primary IP address
	PrimarySubnet *SoftLayer_Network_Subnet `json:"primarySubnet,omitempty"`

	// Guest - The computing instance that this network component exists on.
	Guest *SoftLayer_Virtual_Guest `json:"guest,omitempty"`

	// HighAvailabilityFirewallFlag - <nil>
	HighAvailabilityFirewallFlag bool `json:"highAvailabilityFirewallFlag,omitempty"`

	// IpAddressBindings - The records of all IP addresses bound to a computing instance's network
	// component.
	IpAddressBindings []*SoftLayer_Virtual_Guest_Network_Component_IpAddress `json:"ipAddressBindings,omitempty"`
}

func (softlayer_virtual_guest_network_component *SoftLayer_Virtual_Guest_Network_Component_Extended) String() string {
	return "SoftLayer_Virtual_Guest_Network_Component"
}
