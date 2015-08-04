package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Bandwidth_Version1_Interface - All bandwidth tracking is maintained through the
// switch that the bandwidth is used through. All bandwidth is stored in a "pod" repository. An
// interface links the hardware switch with the pod repository identification number. This is only
// relevant to bandwidth data. It is not common to use this.
type SoftLayer_Network_Bandwidth_Version1_Interface struct {

	// HostId - A interface's host. The host stores the pod number for the bandwidth data.
	HostId int `json:"hostId,omitempty"`

	// NetworkComponentId - no documentation
	NetworkComponentId int `json:"networkComponentId,omitempty"`

	// Host - The host for an interface. This is not to be confused with a SoftLayer hardware
	Host *SoftLayer_Network_Bandwidth_Version1_Host `json:"host,omitempty"`

	// NetworkComponent - no documentation
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent,omitempty"`
}

func (softlayer_network_bandwidth_version1_interface *SoftLayer_Network_Bandwidth_Version1_Interface) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Interface"
}
