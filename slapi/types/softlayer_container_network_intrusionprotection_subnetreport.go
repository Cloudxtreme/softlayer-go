package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_IntrusionProtection_SubnetReport - The IntrusionProtection_SubnetReport
// object is the container that holds the SoftLayer_Container_Network_IntrusionProtection_Event objects
// for a particular subnet, or "All Subnets", whatever the case may be. Subnet, subnet mask, direction,
// and the individual events are returned by this object.
type SoftLayer_Container_Network_IntrusionProtection_SubnetReport struct {

	// Cidr - cidr for this report. If the subnetIpAddress is "All Subnets", this is set to 32 and should
	// be ignored.
	Cidr int `json:"cidr,omitempty"`

	// Direction - Direction of the attack, either 'Inbound' or 'Outbound'
	Direction string `json:"direction,omitempty"`

	// Events - The class SoftLayer_Container_Network_IntrusionProtection_Event objects on this report.
	Events []*SoftLayer_Container_Network_IntrusionProtection_Event `json:"events,omitempty"`

	// SubnetIpAddress - The "target" of this report, could be an IP address, a subnet's network
	// identifier, or "All Subnets"
	SubnetIpAddress string `json:"subnetIpAddress,omitempty"`
}

func (softlayer_container_network_intrusionprotection_subnetreport *SoftLayer_Container_Network_IntrusionProtection_SubnetReport) String() string {
	return "SoftLayer_Container_Network_IntrusionProtection_SubnetReport"
}
