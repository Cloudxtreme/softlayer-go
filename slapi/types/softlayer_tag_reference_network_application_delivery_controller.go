package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Network_Application_Delivery_Controller - <nil>
type SoftLayer_Tag_Reference_Network_Application_Delivery_Controller struct {
}

func (softlayer_tag_reference_network_application_delivery_controller *SoftLayer_Tag_Reference_Network_Application_Delivery_Controller) String() string {
	return "SoftLayer_Tag_Reference_Network_Application_Delivery_Controller"
}

// SoftLayer_Tag_Reference_Network_Application_Delivery_Controller_Extended is SoftLayer_Tag_Reference_Network_Application_Delivery_Controller with all maskable types.
type SoftLayer_Tag_Reference_Network_Application_Delivery_Controller_Extended struct {
	SoftLayer_Tag_Reference_Network_Application_Delivery_Controller

	// Resource - <nil>
	Resource *SoftLayer_Network_Application_Delivery_Controller `json:"resource"`
}

func (softlayer_tag_reference_network_application_delivery_controller *SoftLayer_Tag_Reference_Network_Application_Delivery_Controller_Extended) String() string {
	return "SoftLayer_Tag_Reference_Network_Application_Delivery_Controller"
}
