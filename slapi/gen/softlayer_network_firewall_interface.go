package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Firewall_Interface - The SoftLayer_Network_Firewall_Interface data type contains
// general information relating to a single SoftLayer firewall interface. This is the object which ties
// the firewall context access control list to a firewall. Use the [[SoftLayer Network Firewall
// Template]] service to pull SoftLayer recommended rule set templates. Use the [[SoftLayer Network
// Firewall Update Request]] service to submit a firewall update request.
type SoftLayer_Network_Firewall_Interface struct {
}

// GetObject - getObject returns a SoftLayer_Network_Firewall_Interface object. You can only get
// objects for servers attached to your account that have a network firewall enabled.
func (softlayer_network_firewall_interface *SoftLayer_Network_Firewall_Interface) GetObject() (*SoftLayer_Network_Firewall_Interface, error) {
	var returnValue *SoftLayer_Network_Firewall_Interface
	return returnValue, nil
}
