package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Component - Every piece of hardware running in SoftLayer's datacenters connected
// to the public, private, or management networks (where applicable) have a corresponding network
// component. These network components are modeled by the SoftLayer_Network_Component data type. These
// data types reflect the servers' local ethernet and remote management interfaces.
type SoftLayer_Network_Component struct {

	// MacAddress - A network component's unique MAC address. IPMI-based management network interfaces may
	// not have a MAC address.
	MacAddress string `json:"macAddress"`

	// Port - A network component's port number. Most hardware has more than one network interface. The
	// port property separates these interfaces. Use this in conjunction with the ''name'' property to
	// identify a network component. For instance, the "eth0" interface on a server has the network
	// component name "eth" and port 0.
	Port int `json:"port"`

	// MaxSpeed - A network component's maximum allowed speed, measured in Mbit per second. ''maxSpeed'' is
	// determined by the capabilities of the network interface and the port speed purchased on your
	// SoftLayer server.
	MaxSpeed int `json:"maxSpeed"`

	// DuplexModeId - no documentation
	DuplexModeId string `json:"duplexModeId"`

	// HardwareId - The internal identifier of the hardware that a network component belongs to.
	HardwareId int `json:"hardwareId"`

	// IpmiMacAddress - The MAC address of an IPMI-based management network component.
	IpmiMacAddress string `json:"ipmiMacAddress"`

	// NetworkVlanId - The unique internal id of the network that the port belongs to.
	NetworkVlanId int `json:"networkVlanId"`

	// Id - no documentation
	Id int `json:"id"`

	// PrimaryIpAddress - A network component's primary IP address. IPMI-based management network
	// interfaces may not have an IP address.
	PrimaryIpAddress string `json:"primaryIpAddress"`

	// IpmiIpAddress - The IP address of an IPMI-based management network component.
	IpmiIpAddress string `json:"ipmiIpAddress"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - A network component's short name. For most servers this is the string "eth" for ethernet
	// ports or "mgmt" for remote management ports. Use this in conjunction with the ''port'' property to
	// identify a network component. For instance, the "eth0" interface on a server has the network
	// component name "eth" and port 0.
	Name string `json:"name"`

	// Speed - A network component's speed, measured in Mbit per second.
	Speed int `json:"speed"`

	// Status - A network component's status. This can take one of four possible values: or network
	// components are enabled and in use on a servers. status components have been administratively
	// disabled by SoftLayer accounting or abuse. components have been administratively disabled by you,
	// the user. components only exist on network components that have not been provisioned. You should
	// never see a network interface in state. If you happen to see one please contact SoftLayer support.
	Status string `json:"status"`
}

func (softlayer_network_component *SoftLayer_Network_Component) String() string {
	return "SoftLayer_Network_Component"
}

// SoftLayer_Network_Component_Extended is SoftLayer_Network_Component with all maskable types.
type SoftLayer_Network_Component_Extended struct {
	SoftLayer_Network_Component

	// IpAddressBindingCount - A count of the records of all IP addresses bound to a network component.
	IpAddressBindingCount uint64 `json:"ipAddressBindingCount"`

	// NetworkVlanTrunks - The VLANs that are trunked to this network component.
	NetworkVlanTrunks []*SoftLayer_Network_Component_Network_Vlan_Trunk `json:"networkVlanTrunks"`

	// Subnets - A network component's subnets. A subnet is a group of IP addresses
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// SubnetCount - A count of a network component's subnets. A subnet is a group of IP addresses
	SubnetCount uint64 `json:"subnetCount"`

	// Interface - A hardware switch's interface to the bandwidth pod.
	Interface *SoftLayer_Network_Bandwidth_Version1_Interface `json:"interface"`

	// NetworkComponentFirewall - no documentation
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall"`

	// PrimarySubnet - A network component's subnet for its primary IP address
	PrimarySubnet *SoftLayer_Network_Subnet `json:"primarySubnet"`

	// StorageNetworkFlag - Whether a network component's primary ip address is from a storage network
	// subnet or not.
	StorageNetworkFlag bool `json:"storageNetworkFlag"`

	// RecentCommandCount - A count of the last five reboot/power (rebootDefault, rebootSoft, rebootHard,
	// powerOn, powerOff and powerCycle) commands issued to the server's remote management card.
	RecentCommandCount uint64 `json:"recentCommandCount"`

	// DownlinkComponent - The network component linking this object to a child device
	DownlinkComponent *SoftLayer_Network_Component `json:"downlinkComponent"`

	// NetworkHardware - All network devices in SoftLayer's network hierarchy that this dice is connected
	// to.
	NetworkHardware []*SoftLayer_Hardware `json:"networkHardware"`

	// Router - no documentation
	Router *SoftLayer_Hardware `json:"router"`

	// UplinkDuplexMode - The duplex mode of the uplink network component linking to this object
	UplinkDuplexMode *SoftLayer_Network_Component_Duplex_Mode `json:"uplinkDuplexMode"`

	// HighAvailabilityFirewallFlag - <nil>
	HighAvailabilityFirewallFlag bool `json:"highAvailabilityFirewallFlag"`

	// MetricTrackingObject - The metric tracking object for this network component.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// NetworkVlanTrunkCount - A count of the VLANs that are trunked to this network component.
	NetworkVlanTrunkCount uint64 `json:"networkVlanTrunkCount"`

	// IpAddressCount - no documentation
	IpAddressCount uint64 `json:"ipAddressCount"`

	// NetworkHardwareCount - A count of all network devices in SoftLayer's network hierarchy that this
	// dice is connected to.
	NetworkHardwareCount uint64 `json:"networkHardwareCount"`

	// RedundancyCapableFlag - Indicates whether a server's network component has the ability to be in a
	// redundancy group.
	RedundancyCapableFlag bool `json:"redundancyCapableFlag"`

	// RedundancyEnabledFlag - Indicates whether a server's network component is within an active
	// redundancy group.
	RedundancyEnabledFlag bool `json:"redundancyEnabledFlag"`

	// RemoteManagementUsers - User(s) credentials to issue commands and/or interact with the server's
	// remote management card.
	RemoteManagementUsers []*SoftLayer_Hardware_Component_RemoteManagement_User `json:"remoteManagementUsers"`

	// IpAddresses - <nil>
	IpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"ipAddresses"`

	// PrimaryIpAddressRecord - The primary IPv4 Address record for a network component.
	PrimaryIpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryIpAddressRecord"`

	// NetworkVlan - The that a network component's subnet is associated with.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// IpAddressBindings - The records of all IP addresses bound to a network component.
	IpAddressBindings []*SoftLayer_Network_Component_IpAddress `json:"ipAddressBindings"`

	// NetworkComponentGroup - no documentation
	NetworkComponentGroup *SoftLayer_Network_Component_Group `json:"networkComponentGroup"`

	// PrimaryVersion6IpAddressRecord - The primary IPv6 Address record for a network component.
	PrimaryVersion6IpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryVersion6IpAddressRecord"`

	// RecentCommands - The last five reboot/power (rebootDefault, rebootSoft, rebootHard, powerOn,
	// powerOff and powerCycle) commands issued to the server's remote management card.
	RecentCommands []*SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"recentCommands"`

	// UplinkComponent - The network component linking this object to parent
	UplinkComponent *SoftLayer_Network_Component `json:"uplinkComponent"`

	// DuplexMode - no documentation
	DuplexMode *SoftLayer_Network_Component_Duplex_Mode `json:"duplexMode"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// RemoteManagementUserCount - A count of user(s) credentials to issue commands and/or interact with
	// the server's remote management card.
	RemoteManagementUserCount uint64 `json:"remoteManagementUserCount"`

	// ActiveCommand - Reboot/power (rebootDefault, rebootSoft, rebootHard, powerOn, powerOff and
	// powerCycle) command currently executing by the server's remote management card.
	ActiveCommand *SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"activeCommand"`

	// LastCommand - Last reboot/power (rebootDefault, rebootSoft, rebootHard, powerOn, powerOff and
	// powerCycle) command issued to the server's remote management card.
	LastCommand *SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"lastCommand"`
}

func (softlayer_network_component *SoftLayer_Network_Component_Extended) String() string {
	return "SoftLayer_Network_Component"
}
