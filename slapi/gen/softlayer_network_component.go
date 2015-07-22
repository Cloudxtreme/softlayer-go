package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Component - Every piece of hardware running in SoftLayer's datacenters connected
// to the public, private, or management networks (where applicable) have a corresponding network
// component. These network components are modeled by the SoftLayer_Network_Component data type. These
// data types reflect the servers' local ethernet and remote management interfaces.
type SoftLayer_Network_Component struct {

	// ActiveCommand - Reboot/power (rebootDefault, rebootSoft, rebootHard, powerOn, powerOff and
	// powerCycle) command currently executing by the server's remote management card.
	ActiveCommand *SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"activeCommand"`

	// DownlinkComponent - The network component linking this object to a child device
	DownlinkComponent *SoftLayer_Network_Component `json:"downlinkComponent"`

	// DuplexMode - no documentation
	DuplexMode *SoftLayer_Network_Component_Duplex_Mode `json:"duplexMode"`

	// DuplexModeId - no documentation
	DuplexModeId string `json:"duplexModeId"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - The internal identifier of the hardware that a network component belongs to.
	HardwareId int `json:"hardwareId"`

	// HighAvailabilityFirewallFlag - <nil>
	HighAvailabilityFirewallFlag bool `json:"highAvailabilityFirewallFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// Interface - A hardware switch's interface to the bandwidth pod.
	Interface *SoftLayer_Network_Bandwidth_Version1_Interface `json:"interface"`

	// IpAddressBindingCount - A count of the records of all IP addresses bound to a network component.
	IpAddressBindingCount uint64 `json:"ipAddressBindingCount"`

	// IpAddressBindings - The records of all IP addresses bound to a network component.
	IpAddressBindings []*SoftLayer_Network_Component_IpAddress `json:"ipAddressBindings"`

	// IpAddressCount - no documentation
	IpAddressCount uint64 `json:"ipAddressCount"`

	// IpAddresses - <nil>
	IpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"ipAddresses"`

	// IpmiIpAddress - The IP address of an IPMI-based management network component.
	IpmiIpAddress string `json:"ipmiIpAddress"`

	// IpmiMacAddress - The MAC address of an IPMI-based management network component.
	IpmiMacAddress string `json:"ipmiMacAddress"`

	// LastCommand - Last reboot/power (rebootDefault, rebootSoft, rebootHard, powerOn, powerOff and
	// powerCycle) command issued to the server's remote management card.
	LastCommand *SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"lastCommand"`

	// MacAddress - A network component's unique MAC address. IPMI-based management network interfaces may
	// not have a MAC address.
	MacAddress string `json:"macAddress"`

	// MaxSpeed - A network component's maximum allowed speed, measured in Mbit per second. ''maxSpeed'' is
	// determined by the capabilities of the network interface and the port speed purchased on your
	// SoftLayer server.
	MaxSpeed int `json:"maxSpeed"`

	// MetricTrackingObject - The metric tracking object for this network component.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - A network component's short name. For most servers this is the string "eth" for ethernet
	// ports or "mgmt" for remote management ports. Use this in conjunction with the ''port'' property to
	// identify a network component. For instance, the "eth0" interface on a server has the network
	// component name "eth" and port 0.
	Name string `json:"name"`

	// NetworkComponentFirewall - no documentation
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall"`

	// NetworkComponentGroup - no documentation
	NetworkComponentGroup *SoftLayer_Network_Component_Group `json:"networkComponentGroup"`

	// NetworkHardware - All network devices in SoftLayer's network hierarchy that this dice is connected
	// to.
	NetworkHardware []*SoftLayer_Hardware `json:"networkHardware"`

	// NetworkHardwareCount - A count of all network devices in SoftLayer's network hierarchy that this
	// dice is connected to.
	NetworkHardwareCount uint64 `json:"networkHardwareCount"`

	// NetworkVlan - The that a network component's subnet is associated with.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// NetworkVlanId - The unique internal id of the network that the port belongs to.
	NetworkVlanId int `json:"networkVlanId"`

	// NetworkVlanTrunkCount - A count of the VLANs that are trunked to this network component.
	NetworkVlanTrunkCount uint64 `json:"networkVlanTrunkCount"`

	// NetworkVlanTrunks - The VLANs that are trunked to this network component.
	NetworkVlanTrunks []*SoftLayer_Network_Component_Network_Vlan_Trunk `json:"networkVlanTrunks"`

	// Port - A network component's port number. Most hardware has more than one network interface. The
	// port property separates these interfaces. Use this in conjunction with the ''name'' property to
	// identify a network component. For instance, the "eth0" interface on a server has the network
	// component name "eth" and port 0.
	Port int `json:"port"`

	// PrimaryIpAddress - A network component's primary IP address. IPMI-based management network
	// interfaces may not have an IP address.
	PrimaryIpAddress string `json:"primaryIpAddress"`

	// PrimaryIpAddressRecord - The primary IPv4 Address record for a network component.
	PrimaryIpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryIpAddressRecord"`

	// PrimarySubnet - A network component's subnet for its primary IP address
	PrimarySubnet *SoftLayer_Network_Subnet `json:"primarySubnet"`

	// PrimaryVersion6IpAddressRecord - The primary IPv6 Address record for a network component.
	PrimaryVersion6IpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"primaryVersion6IpAddressRecord"`

	// RecentCommandCount - A count of the last five reboot/power (rebootDefault, rebootSoft, rebootHard,
	// powerOn, powerOff and powerCycle) commands issued to the server's remote management card.
	RecentCommandCount uint64 `json:"recentCommandCount"`

	// RecentCommands - The last five reboot/power (rebootDefault, rebootSoft, rebootHard, powerOn,
	// powerOff and powerCycle) commands issued to the server's remote management card.
	RecentCommands []*SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"recentCommands"`

	// RedundancyCapableFlag - Indicates whether a server's network component has the ability to be in a
	// redundancy group.
	RedundancyCapableFlag bool `json:"redundancyCapableFlag"`

	// RedundancyEnabledFlag - Indicates whether a server's network component is within an active
	// redundancy group.
	RedundancyEnabledFlag bool `json:"redundancyEnabledFlag"`

	// RemoteManagementUserCount - A count of user(s) credentials to issue commands and/or interact with
	// the server's remote management card.
	RemoteManagementUserCount uint64 `json:"remoteManagementUserCount"`

	// RemoteManagementUsers - User(s) credentials to issue commands and/or interact with the server's
	// remote management card.
	RemoteManagementUsers []*SoftLayer_Hardware_Component_RemoteManagement_User `json:"remoteManagementUsers"`

	// Router - no documentation
	Router *SoftLayer_Hardware `json:"router"`

	// Speed - A network component's speed, measured in Mbit per second.
	Speed int `json:"speed"`

	// Status - A network component's status. This can take one of four possible values: or network
	// components are enabled and in use on a servers. status components have been administratively
	// disabled by SoftLayer accounting or abuse. components have been administratively disabled by you,
	// the user. components only exist on network components that have not been provisioned. You should
	// never see a network interface in state. If you happen to see one please contact SoftLayer support.
	Status string `json:"status"`

	// StorageNetworkFlag - Whether a network component's primary ip address is from a storage network
	// subnet or not.
	StorageNetworkFlag bool `json:"storageNetworkFlag"`

	// SubnetCount - A count of a network component's subnets. A subnet is a group of IP addresses
	SubnetCount uint64 `json:"subnetCount"`

	// Subnets - A network component's subnets. A subnet is a group of IP addresses
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// UplinkComponent - The network component linking this object to parent
	UplinkComponent *SoftLayer_Network_Component `json:"uplinkComponent"`

	// UplinkDuplexMode - The duplex mode of the uplink network component linking to this object
	UplinkDuplexMode *SoftLayer_Network_Component_Duplex_Mode `json:"uplinkDuplexMode"`
}

// AddNetworkVlanTrunks - Add VLANs as trunks to a network component. The VLANs given must be assigned
// to your account, and on the router to which this network component is connected. The current native
// (networkVlanId/networkVlan) cannot be added as a trunk. This method should be called on a network
// component attached directly to customer assigned hardware, though all trunking operations will occur
// on the uplinkComponent. A current list of trunks for a network component on a customer server can be
// found at 'uplinkComponent->networkVlanTrunks'. This method returns an array of
// SoftLayer_Network_Vlans which were added as trunks. Any requested trunks which are already trunked
// will be silently ignored, and will not be returned. Configuration of network hardware is done
// asynchronously, do not depend on the return of this call as an indication that the newly trunked
// VLANs will be accessible.
func (softlayer_network_component *SoftLayer_Network_Component) AddNetworkVlanTrunks(networkVlans []SoftLayer_Network_Vlan) ([]*SoftLayer_Network_Vlan, error) {
	var returnValue []*SoftLayer_Network_Vlan
	return returnValue, nil
}

// ClearNetworkVlanTrunks - This method will remove all VLANs trunked to this network component. The
// native (networkVlanId/networkVlan) will remain active, and cannot be removed via the Returns a list
// of SoftLayer_Network_Vlan objects for which the trunks were removed.
func (softlayer_network_component *SoftLayer_Network_Component) ClearNetworkVlanTrunks() ([]*SoftLayer_Network_Vlan, error) {
	var returnValue []*SoftLayer_Network_Vlan
	return returnValue, nil
}

// GetCustomBandwidthDataByDate - no documentation
func (softlayer_network_component *SoftLayer_Network_Component) GetCustomBandwidthDataByDate(graphData SoftLayer_Container_Graph) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_component *SoftLayer_Network_Component) GetObject() (*SoftLayer_Network_Component, error) {
	var returnValue *SoftLayer_Network_Component
	return returnValue, nil
}

// GetPortStatistics - Retrieve various network statistics. The network statistics are retrieved from
// the network device using snmpget. Below is a list of statistics retrieved: * Administrative Status *
// Operational Status * Maximum Transmission Unit * In Octets * Out Octets * In Unicast Packets * Out
// Unicast Packets * In Multicast Packets * Out Multicast Packets
func (softlayer_network_component *SoftLayer_Network_Component) GetPortStatistics() (*SoftLayer_Container_Network_Port_Statistic, error) {
	var returnValue *SoftLayer_Container_Network_Port_Statistic
	return returnValue, nil
}

// RemoveNetworkVlanTrunks - Remove one or more VLANs currently attached to the uplinkComponent of this
// networkComponent. The VLANs given must be assigned to your account, and on the router the network
// component is connected to. If any VLANs not currently trunked are given, they will be silently
// ignored. This method should be called on a network component attached directly to customer assigned
// hardware, though all trunking operations will occur on the uplinkComponent. A current list of trunks
// for a network component on a customer server can be found at 'uplinkComponent->networkVlanTrunks'.
// Configuration of network hardware is done asynchronously, do not depend on the return of this call
// as an indication that the removed VLANs will be inaccessible.
func (softlayer_network_component *SoftLayer_Network_Component) RemoveNetworkVlanTrunks(networkVlans []SoftLayer_Network_Vlan) ([]*SoftLayer_Network_Vlan, error) {
	var returnValue []*SoftLayer_Network_Vlan
	return returnValue, nil
}
