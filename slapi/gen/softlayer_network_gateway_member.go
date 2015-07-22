package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Gateway_Member - <nil>
type SoftLayer_Network_Gateway_Member struct {

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - The internal identifier of the hardware for this member.
	HardwareId int `json:"hardwareId"`

	// Id - no documentation
	Id int `json:"id"`

	// NetworkGateway - no documentation
	NetworkGateway *SoftLayer_Network_Gateway `json:"networkGateway"`

	// NetworkGatewayId - The internal identifier of the gateway this member belongs to.
	NetworkGatewayId int `json:"networkGatewayId"`

	// Priority - The priority for this gateway member. This is set internally and cannot be provided on
	// create.
	Priority int `json:"priority"`
}

// CreateObject - Create a new hardware member on the gateway. This also asynchronously sets up the
// network for this member. Progress of this process can be monitored via the gateway status. All
// members created with this object must have no VLANs attached.
func (softlayer_network_gateway_member *SoftLayer_Network_Gateway_Member) CreateObject(templateObject SoftLayer_Network_Gateway_Member) (*SoftLayer_Network_Gateway_Member, error) {
	var returnValue *SoftLayer_Network_Gateway_Member
	return returnValue, nil
}

// CreateObjects - Create multiple new hardware members on the gateway. This also asynchronously sets
// up the network for the members. Progress of this process can be monitored via the gateway status.
// All members created with this object must have no VLANs attached.
func (softlayer_network_gateway_member *SoftLayer_Network_Gateway_Member) CreateObjects(templateObjects []SoftLayer_Network_Gateway_Member) ([]*SoftLayer_Network_Gateway_Member, error) {
	var returnValue []*SoftLayer_Network_Gateway_Member
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_gateway_member *SoftLayer_Network_Gateway_Member) GetObject() (*SoftLayer_Network_Gateway_Member, error) {
	var returnValue *SoftLayer_Network_Gateway_Member
	return returnValue, nil
}
