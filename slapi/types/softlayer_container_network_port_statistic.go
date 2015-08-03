package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Port_Statistic - <nil>
type SoftLayer_Container_Network_Port_Statistic struct {

	// Speed - <nil>
	Speed uint64 `json:"speed"`

	// AdministrativeStatus - <nil>
	AdministrativeStatus int `json:"administrativeStatus"`

	// InOctets - <nil>
	InOctets uint64 `json:"inOctets"`

	// InUnicastPackets - <nil>
	InUnicastPackets uint64 `json:"inUnicastPackets"`

	// OutErrorPackets - <nil>
	OutErrorPackets uint64 `json:"outErrorPackets"`

	// OutDiscardPackets - <nil>
	OutDiscardPackets uint64 `json:"outDiscardPackets"`

	// OutOctets - <nil>
	OutOctets uint64 `json:"outOctets"`

	// OutUnicastPackets - <nil>
	OutUnicastPackets uint64 `json:"outUnicastPackets"`

	// PortDuplex - <nil>
	PortDuplex uint64 `json:"portDuplex"`

	// InDiscardPackets - <nil>
	InDiscardPackets uint64 `json:"inDiscardPackets"`

	// InErrorPackets - <nil>
	InErrorPackets uint64 `json:"inErrorPackets"`

	// MaximumTransmissionUnit - <nil>
	MaximumTransmissionUnit uint64 `json:"maximumTransmissionUnit"`

	// OperationalStatus - <nil>
	OperationalStatus int `json:"operationalStatus"`
}

func (softlayer_container_network_port_statistic *SoftLayer_Container_Network_Port_Statistic) String() string {
	return "SoftLayer_Container_Network_Port_Statistic"
}
