package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource_Attribute - <nil>
type SoftLayer_Network_Service_Resource_Attribute struct {

	// Value - <nil>
	Value string `json:"value"`
}

// SoftLayer_Network_Service_Resource_Attribute_Extended is SoftLayer_Network_Service_Resource_Attribute with all maskable types.
type SoftLayer_Network_Service_Resource_Attribute_Extended struct {
	SoftLayer_Network_Service_Resource_Attribute

	// AttributeType - <nil>
	AttributeType *SoftLayer_Network_Service_Resource_Attribute_Type `json:"attributeType"`

	// ServiceResource - <nil>
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource"`
}

func (softlayer_network_service_resource_attribute *SoftLayer_Network_Service_Resource_Attribute) String() string {
	return "SoftLayer_Network_Service_Resource_Attribute"
}
