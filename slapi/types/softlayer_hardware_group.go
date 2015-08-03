package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Group - <nil>
type SoftLayer_Hardware_Group struct {

	// Domain - <nil>
	Domain string `json:"domain"`

	// Hostname - <nil>
	Hostname string `json:"hostname"`

	// Id - <nil>
	Id int `json:"id"`
}

// SoftLayer_Hardware_Group_Extended is SoftLayer_Hardware_Group with all maskable types.
type SoftLayer_Hardware_Group_Extended struct {
	SoftLayer_Hardware_Group

	// NetworkMonitorAttachedDownVirtualGuests - Virtual guests that are attached downstream to a hardware
	// that have failed monitoring
	NetworkMonitorAttachedDownVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorAttachedDownVirtualGuests"`

	// DownlinkServerCount - A count of all servers attached to a network hardware.
	DownlinkServerCount uint64 `json:"downlinkServerCount"`

	// NetworkMonitorAttachedDownVirtualGuestCount - A count of virtual guests that are attached downstream
	// to a hardware that have failed monitoring
	NetworkMonitorAttachedDownVirtualGuestCount uint64 `json:"networkMonitorAttachedDownVirtualGuestCount"`

	// DownlinkVirtualGuests - no documentation
	DownlinkVirtualGuests []*SoftLayer_Virtual_Guest `json:"downlinkVirtualGuests"`

	// DownstreamNetworkHardwareWithIncidents - All network hardware with monitoring warnings or errors
	// downstream from this hardware.
	DownstreamNetworkHardwareWithIncidents []*SoftLayer_Hardware `json:"downstreamNetworkHardwareWithIncidents"`

	// NetworkMonitorAttachedDownHardwareCount - A count of all servers attached downstream to a hardware
	// that have failed monitoring
	NetworkMonitorAttachedDownHardwareCount uint64 `json:"networkMonitorAttachedDownHardwareCount"`

	// DownlinkVirtualGuestCount - A count of all virtual guests attached to a network hardware.
	DownlinkVirtualGuestCount uint64 `json:"downlinkVirtualGuestCount"`

	// DownstreamNetworkHardwareCount - A count of all network hardware downstream from this hardware.
	DownstreamNetworkHardwareCount uint64 `json:"downstreamNetworkHardwareCount"`

	// DownstreamNetworkHardwareWithIncidentCount - A count of all network hardware with monitoring
	// warnings or errors downstream from this hardware.
	DownstreamNetworkHardwareWithIncidentCount uint64 `json:"downstreamNetworkHardwareWithIncidentCount"`

	// HardwareChassis - no documentation
	HardwareChassis *SoftLayer_Hardware_Chassis `json:"hardwareChassis"`

	// NetworkStatus - The value of a hardware's network status attribute.
	NetworkStatus string `json:"networkStatus"`

	// NetworkMonitorAttachedDownHardware - All servers attached downstream to a hardware that have failed
	// monitoring
	NetworkMonitorAttachedDownHardware []*SoftLayer_Hardware `json:"networkMonitorAttachedDownHardware"`

	// DownlinkServers - no documentation
	DownlinkServers []*SoftLayer_Hardware `json:"downlinkServers"`

	// DownstreamNetworkHardware - All network hardware downstream from this hardware.
	DownstreamNetworkHardware []*SoftLayer_Hardware `json:"downstreamNetworkHardware"`
}

func (softlayer_hardware_group *SoftLayer_Hardware_Group) String() string {
	return "SoftLayer_Hardware_Group"
}
