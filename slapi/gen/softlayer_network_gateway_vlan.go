package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Gateway_Vlan - <nil>
type SoftLayer_Network_Gateway_Vlan struct {

	// BypassFlag - If true, this is bypassed. If false, it is routed through the gateway.
	BypassFlag bool `json:"bypassFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// NetworkGateway - no documentation
	NetworkGateway *SoftLayer_Network_Gateway `json:"networkGateway"`

	// NetworkGatewayId - The internal identifier of the gateway this is attached to.
	NetworkGatewayId int `json:"networkGatewayId"`

	// NetworkVlan - no documentation
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// NetworkVlanId - no documentation
	NetworkVlanId int `json:"networkVlanId"`
}

// Bypass - Start the asynchronous process to bypass/unroute the from this gateway.
func (softlayer_network_gateway_vlan *SoftLayer_Network_Gateway_Vlan) Bypass() error {
	return nil
}

// CreateObject - Create a new attachment. If the bypassFlag is false, this will also create an
// asynchronous process to route the through the gateway.
func (softlayer_network_gateway_vlan *SoftLayer_Network_Gateway_Vlan) CreateObject(templateObject SoftLayer_Network_Gateway_Vlan) (*SoftLayer_Network_Gateway_Vlan, error) {
	var returnValue *SoftLayer_Network_Gateway_Vlan
	return returnValue, nil
}

// CreateObjects - Create multiple new attachments. If the bypassFlag is false, this will also create
// an asynchronous process to route the VLANs through the gateway.
func (softlayer_network_gateway_vlan *SoftLayer_Network_Gateway_Vlan) CreateObjects(templateObjects []SoftLayer_Network_Gateway_Vlan) ([]*SoftLayer_Network_Gateway_Vlan, error) {
	var returnValue []*SoftLayer_Network_Gateway_Vlan
	return returnValue, nil
}

// DeleteObject - Start the asynchronous process to detach this VLANs from the gateway.
func (softlayer_network_gateway_vlan *SoftLayer_Network_Gateway_Vlan) DeleteObject() error {
	return nil
}

// DeleteObjects - Detach several VLANs. This will not detach them right away, but rather start an
// asynchronous process to detach.
func (softlayer_network_gateway_vlan *SoftLayer_Network_Gateway_Vlan) DeleteObjects(templateObjects []SoftLayer_Network_Gateway_Vlan) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_gateway_vlan *SoftLayer_Network_Gateway_Vlan) GetObject() (*SoftLayer_Network_Gateway_Vlan, error) {
	var returnValue *SoftLayer_Network_Gateway_Vlan
	return returnValue, nil
}

// Unbypass - Start the asynchronous process to route the to this gateway.
func (softlayer_network_gateway_vlan *SoftLayer_Network_Gateway_Vlan) Unbypass() error {
	return nil
}
