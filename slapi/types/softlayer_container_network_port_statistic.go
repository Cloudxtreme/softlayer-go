package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Port_Statistic - <nil>
type SoftLayer_Container_Network_Port_Statistic struct {

	// InErrorPackets - <nil>
	InErrorPackets uint64 `json:"inErrorPackets,omitempty"`

	// InUnicastPackets - <nil>
	InUnicastPackets uint64 `json:"inUnicastPackets,omitempty"`

	// OperationalStatus - <nil>
	OperationalStatus int `json:"operationalStatus,omitempty"`

	// OutDiscardPackets - <nil>
	OutDiscardPackets uint64 `json:"outDiscardPackets,omitempty"`

	// OutErrorPackets - <nil>
	OutErrorPackets uint64 `json:"outErrorPackets,omitempty"`

	// PortDuplex - <nil>
	PortDuplex uint64 `json:"portDuplex,omitempty"`

	// Speed - <nil>
	Speed uint64 `json:"speed,omitempty"`

	// AdministrativeStatus - <nil>
	AdministrativeStatus int `json:"administrativeStatus,omitempty"`

	// InDiscardPackets - <nil>
	InDiscardPackets uint64 `json:"inDiscardPackets,omitempty"`

	// InOctets - <nil>
	InOctets uint64 `json:"inOctets,omitempty"`

	// MaximumTransmissionUnit - <nil>
	MaximumTransmissionUnit uint64 `json:"maximumTransmissionUnit,omitempty"`

	// OutOctets - <nil>
	OutOctets uint64 `json:"outOctets,omitempty"`

	// OutUnicastPackets - <nil>
	OutUnicastPackets uint64 `json:"outUnicastPackets,omitempty"`
}

func (softlayer_container_network_port_statistic *SoftLayer_Container_Network_Port_Statistic) String() string {
	return "SoftLayer_Container_Network_Port_Statistic"
}
