package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol
// - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol struct {

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// VirtualIpAddressId - <nil>
	VirtualIpAddressId int `json:"virtualIpAddressId"`

	// Id - <nil>
	Id int `json:"id"`
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol

	// VirtualIpAddress - <nil>
	VirtualIpAddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer `json:"virtualIpAddress"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress_securetransportprotocol *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol"
}
