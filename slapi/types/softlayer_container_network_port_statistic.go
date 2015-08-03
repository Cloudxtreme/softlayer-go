package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Port_Statistic - <nil>
type SoftLayer_Container_Network_Port_Statistic struct {

	// OperationalStatus - <nil>
	OperationalStatus int `json:"operationalStatus"`

	// OutOctets - <nil>
	OutOctets uint64 `json:"outOctets"`

	// OutUnicastPackets - <nil>
	OutUnicastPackets uint64 `json:"outUnicastPackets"`

	// InDiscardPackets - <nil>
	InDiscardPackets uint64 `json:"inDiscardPackets"`

	// InOctets - <nil>
	InOctets uint64 `json:"inOctets"`

	// InUnicastPackets - <nil>
	InUnicastPackets uint64 `json:"inUnicastPackets"`

	// MaximumTransmissionUnit - <nil>
	MaximumTransmissionUnit uint64 `json:"maximumTransmissionUnit"`

	// PortDuplex - <nil>
	PortDuplex uint64 `json:"portDuplex"`

	// Speed - <nil>
	Speed uint64 `json:"speed"`

	// AdministrativeStatus - <nil>
	AdministrativeStatus int `json:"administrativeStatus"`

	// InErrorPackets - <nil>
	InErrorPackets uint64 `json:"inErrorPackets"`

	// OutDiscardPackets - <nil>
	OutDiscardPackets uint64 `json:"outDiscardPackets"`

	// OutErrorPackets - <nil>
	OutErrorPackets uint64 `json:"outErrorPackets"`
}

func (softlayer_container_network_port_statistic *SoftLayer_Container_Network_Port_Statistic) String() string {
	return "SoftLayer_Container_Network_Port_Statistic"
}
