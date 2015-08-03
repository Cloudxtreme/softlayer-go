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

	// Id - no documentation
	Id int `json:"id"`

	// IpmiIpAddress - The IP address of an IPMI-based management network component.
	IpmiIpAddress string `json:"ipmiIpAddress"`

	// IpmiMacAddress - The MAC address of an IPMI-based management network component.
	IpmiMacAddress string `json:"ipmiMacAddress"`

	// HardwareId - The internal identifier of the hardware that a network component belongs to.
	HardwareId int `json:"hardwareId"`

	// MaxSpeed - A network component's maximum allowed speed, measured in Mbit per second. ''maxSpeed'' is
	// determined by the capabilities of the network interface and the port speed purchased on your
	// SoftLayer server.
	MaxSpeed int `json:"maxSpeed"`

	// Name - A network component's short name. For most servers this is the string "eth" for ethernet
	// ports or "mgmt" for remote management ports. Use this in conjunction with the ''port'' property to
	// identify a network component. For instance, the "eth0" interface on a server has the network
	// component name "eth" and port 0.
	Name string `json:"name"`

	// NetworkVlanId - The unique internal id of the network that the port belongs to.
	NetworkVlanId int `json:"networkVlanId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Status - A network component's status. This can take one of four possible values: or network
	// components are enabled and in use on a servers. status components have been administratively
	// disabled by SoftLayer accounting or abuse. components have been administratively disabled by you,
	// the user. components only exist on network components that have not been provisioned. You should
	// never see a network interface in state. If you happen to see one please contact SoftLayer support.
	Status string `json:"status"`

	// DuplexModeId - no documentation
	DuplexModeId string `json:"duplexModeId"`

	// Port - A network component's port number. Most hardware has more than one network interface. The
	// port property separates these interfaces. Use this in conjunction with the ''name'' property to
	// identify a network component. For instance, the "eth0" interface on a server has the network
	// component name "eth" and port 0.
	Port int `json:"port"`

	// MacAddress - A network component's unique MAC address. IPMI-based management network interfaces may
	// not have a MAC address.
	MacAddress string `json:"macAddress"`

	// Speed - A network component's speed, measured in Mbit per second.
	Speed int `json:"speed"`

	// PrimaryIpAddress - A network component's primary IP address. IPMI-based management network
	// interfaces may not have an IP address.
	PrimaryIpAddress string `json:"primaryIpAddress"`
}

// SoftLayer_Network_Component_Extended is SoftLayer_Network_Component with all maskable types.
type SoftLayer_Network_Component_Extended struct {
	SoftLayer_Network_Component

	// NetworkVlanTrunks - The VLANs that are trunked to this network component.
	NetworkVlanTrunks []*SoftLayer_Network_Component_Network_Vlan_Trunk `json:"networkVlanTrunks"`

	// RecentCommands - The last five reboot/power (rebootDefault, rebootSoft, rebootHard, powerOn,
	// powerOff and powerCycle) commands issued to the server's remote management card.
	RecentCommands []*SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"recentCommands"`

	// HighAvailabilityFirewallFlag - <nil>
	HighAvailabilityFirewallFlag bool `json:"highAvailabilityFirewallFlag"`

	// NetworkComponentFirewall - no documentation
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall"`

	// PrimaryIpAddressRecord - The primary IPv4 Address record for a network component.
	PrimaryIpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryIpAddressRecord"`

	// PrimarySubnet - A network component's subnet for its primary IP address
	PrimarySubnet *SoftLayer_Network_Subnet `json:"primarySubnet"`

	// PrimaryVersion6IpAddressRecord - The primary IPv6 Address record for a network component.
	PrimaryVersion6IpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryVersion6IpAddressRecord"`

	// IpAddressBindingCount - A count of the records of all IP addresses bound to a network component.
	IpAddressBindingCount uint64 `json:"ipAddressBindingCount"`

	// RemoteManagementUserCount - A count of user(s) credentials to issue commands and/or interact with
	// the server's remote management card.
	RemoteManagementUserCount uint64 `json:"remoteManagementUserCount"`

	// RedundancyCapableFlag - Indicates whether a server's network component has the ability to be in a
	// redundancy group.
	RedundancyCapableFlag bool `json:"redundancyCapableFlag"`

	// UplinkDuplexMode - The duplex mode of the uplink network component linking to this object
	UplinkDuplexMode *SoftLayer_Network_Component_Duplex_Mode `json:"uplinkDuplexMode"`

	// IpAddresses - <nil>
	IpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"ipAddresses"`

	// NetworkVlan - The that a network component's subnet is associated with.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// RemoteManagementUsers - User(s) credentials to issue commands and/or interact with the server's
	// remote management card.
	RemoteManagementUsers []*SoftLayer_Hardware_Component_RemoteManagement_User `json:"remoteManagementUsers"`

	// Router - no documentation
	Router *SoftLayer_Hardware `json:"router"`

	// RecentCommandCount - A count of the last five reboot/power (rebootDefault, rebootSoft, rebootHard,
	// powerOn, powerOff and powerCycle) commands issued to the server's remote management card.
	RecentCommandCount uint64 `json:"recentCommandCount"`

	// NetworkHardwareCount - A count of all network devices in SoftLayer's network hierarchy that this
	// dice is connected to.
	NetworkHardwareCount uint64 `json:"networkHardwareCount"`

	// NetworkVlanTrunkCount - A count of the VLANs that are trunked to this network component.
	NetworkVlanTrunkCount uint64 `json:"networkVlanTrunkCount"`

	// DownlinkComponent - The network component linking this object to a child device
	DownlinkComponent *SoftLayer_Network_Component `json:"downlinkComponent"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// MetricTrackingObject - The metric tracking object for this network component.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// RedundancyEnabledFlag - Indicates whether a server's network component is within an active
	// redundancy group.
	RedundancyEnabledFlag bool `json:"redundancyEnabledFlag"`

	// DuplexMode - no documentation
	DuplexMode *SoftLayer_Network_Component_Duplex_Mode `json:"duplexMode"`

	// NetworkComponentGroup - no documentation
	NetworkComponentGroup *SoftLayer_Network_Component_Group `json:"networkComponentGroup"`

	// Subnets - A network component's subnets. A subnet is a group of IP addresses
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// IpAddressCount - no documentation
	IpAddressCount uint64 `json:"ipAddressCount"`

	// Interface - A hardware switch's interface to the bandwidth pod.
	Interface *SoftLayer_Network_Bandwidth_Version1_Interface `json:"interface"`

	// LastCommand - Last reboot/power (rebootDefault, rebootSoft, rebootHard, powerOn, powerOff and
	// powerCycle) command issued to the server's remote management card.
	LastCommand *SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"lastCommand"`

	// NetworkHardware - All network devices in SoftLayer's network hierarchy that this dice is connected
	// to.
	NetworkHardware []*SoftLayer_Hardware `json:"networkHardware"`

	// StorageNetworkFlag - Whether a network component's primary ip address is from a storage network
	// subnet or not.
	StorageNetworkFlag bool `json:"storageNetworkFlag"`

	// UplinkComponent - The network component linking this object to parent
	UplinkComponent *SoftLayer_Network_Component `json:"uplinkComponent"`

	// ActiveCommand - Reboot/power (rebootDefault, rebootSoft, rebootHard, powerOn, powerOff and
	// powerCycle) command currently executing by the server's remote management card.
	ActiveCommand *SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"activeCommand"`

	// IpAddressBindings - The records of all IP addresses bound to a network component.
	IpAddressBindings []*SoftLayer_Network_Component_IpAddress `json:"ipAddressBindings"`

	// SubnetCount - A count of a network component's subnets. A subnet is a group of IP addresses
	SubnetCount uint64 `json:"subnetCount"`
}

func (softlayer_network_component *SoftLayer_Network_Component) String() string {
	return "SoftLayer_Network_Component"
}
