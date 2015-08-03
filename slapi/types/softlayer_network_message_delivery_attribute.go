package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Message_Delivery_Attribute - <nil>
type SoftLayer_Network_Message_Delivery_Attribute struct {

	// Value - <nil>
	Value string `json:"value"`
}

func (softlayer_network_message_delivery_attribute *SoftLayer_Network_Message_Delivery_Attribute) String() string {
	return "SoftLayer_Network_Message_Delivery_Attribute"
}

// SoftLayer_Network_Message_Delivery_Attribute_Extended is SoftLayer_Network_Message_Delivery_Attribute with all maskable types.
type SoftLayer_Network_Message_Delivery_Attribute_Extended struct {
	SoftLayer_Network_Message_Delivery_Attribute

	// NetworkMessageDelivery - <nil>
	NetworkMessageDelivery *SoftLayer_Network_Message_Delivery `json:"networkMessageDelivery"`
}

func (softlayer_network_message_delivery_attribute *SoftLayer_Network_Message_Delivery_Attribute_Extended) String() string {
	return "SoftLayer_Network_Message_Delivery_Attribute"
}
