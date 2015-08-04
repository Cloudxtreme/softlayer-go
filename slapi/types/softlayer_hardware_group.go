package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Group - <nil>
type SoftLayer_Hardware_Group struct {

	// Hostname - <nil>
	Hostname string `json:"hostname,omitempty"`

	// Domain - <nil>
	Domain string `json:"domain,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`
}

func (softlayer_hardware_group *SoftLayer_Hardware_Group) String() string {
	return "SoftLayer_Hardware_Group"
}

// SoftLayer_Hardware_Group_Extended is SoftLayer_Hardware_Group with all maskable types.
type SoftLayer_Hardware_Group_Extended struct {
	SoftLayer_Hardware_Group

	// NetworkMonitorAttachedDownHardware - All servers attached downstream to a hardware that have failed
	// monitoring
	NetworkMonitorAttachedDownHardware []*SoftLayer_Hardware `json:"networkMonitorAttachedDownHardware,omitempty"`

	// NetworkMonitorAttachedDownVirtualGuestCount - A count of virtual guests that are attached downstream
	// to a hardware that have failed monitoring
	NetworkMonitorAttachedDownVirtualGuestCount uint64 `json:"networkMonitorAttachedDownVirtualGuestCount,omitempty"`

	// DownstreamNetworkHardwareWithIncidentCount - A count of all network hardware with monitoring
	// warnings or errors downstream from this hardware.
	DownstreamNetworkHardwareWithIncidentCount uint64 `json:"downstreamNetworkHardwareWithIncidentCount,omitempty"`

	// DownlinkServers - no documentation
	DownlinkServers []*SoftLayer_Hardware `json:"downlinkServers,omitempty"`

	// DownlinkVirtualGuests - no documentation
	DownlinkVirtualGuests []*SoftLayer_Virtual_Guest `json:"downlinkVirtualGuests,omitempty"`

	// HardwareChassis - no documentation
	HardwareChassis *SoftLayer_Hardware_Chassis `json:"hardwareChassis,omitempty"`

	// NetworkMonitorAttachedDownVirtualGuests - Virtual guests that are attached downstream to a hardware
	// that have failed monitoring
	NetworkMonitorAttachedDownVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorAttachedDownVirtualGuests,omitempty"`

	// DownlinkServerCount - A count of all servers attached to a network hardware.
	DownlinkServerCount uint64 `json:"downlinkServerCount,omitempty"`

	// DownstreamNetworkHardwareCount - A count of all network hardware downstream from this hardware.
	DownstreamNetworkHardwareCount uint64 `json:"downstreamNetworkHardwareCount,omitempty"`

	// DownstreamNetworkHardwareWithIncidents - All network hardware with monitoring warnings or errors
	// downstream from this hardware.
	DownstreamNetworkHardwareWithIncidents []*SoftLayer_Hardware `json:"downstreamNetworkHardwareWithIncidents,omitempty"`

	// DownlinkVirtualGuestCount - A count of all virtual guests attached to a network hardware.
	DownlinkVirtualGuestCount uint64 `json:"downlinkVirtualGuestCount,omitempty"`

	// DownstreamNetworkHardware - All network hardware downstream from this hardware.
	DownstreamNetworkHardware []*SoftLayer_Hardware `json:"downstreamNetworkHardware,omitempty"`

	// NetworkStatus - The value of a hardware's network status attribute.
	NetworkStatus string `json:"networkStatus,omitempty"`

	// NetworkMonitorAttachedDownHardwareCount - A count of all servers attached downstream to a hardware
	// that have failed monitoring
	NetworkMonitorAttachedDownHardwareCount uint64 `json:"networkMonitorAttachedDownHardwareCount,omitempty"`
}

func (softlayer_hardware_group *SoftLayer_Hardware_Group_Extended) String() string {
	return "SoftLayer_Hardware_Group"
}
