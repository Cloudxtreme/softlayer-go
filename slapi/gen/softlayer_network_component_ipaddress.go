package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_IpAddress - The SoftLayer_Network_Component_IpAddress data type contains
// general information relating to the binding of a single network component to a single SoftLayer IP
// address.
type SoftLayer_Network_Component_IpAddress struct {

	// IpAddress - The IP address associated with this object's network component.
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress"`

	// NetworkComponent - The network component associated with this object's IP address.
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`
}
